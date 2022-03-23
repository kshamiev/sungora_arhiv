package stpg

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"runtime/debug"
	"sync"

	"sungora/lib/storage"

	"contrib.go.opencensus.io/integrations/ocsql"
	// драйвер работы с БД
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db  *sqlx.DB
	pqs map[string]*sqlx.Stmt
	mu  sync.RWMutex
}

var instance *Storage

func InitConnect(cfg *Config) error {
	if cfg != nil {
		config = cfg
	}
	if config == nil {
		return errors.New("config is empty")
	}
	cfg.Postgres = fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.User,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		cfg.Dbname,
	)

	var err error
	driverName := "pgx"
	if config.OcSQLTrace {
		driverName, err = ocsql.Register(
			driverName,
			ocsql.WithOptions(
				ocsql.TraceOptions{
					AllowRoot:    true,
					Ping:         true,
					RowsNext:     false,
					RowsClose:    false,
					RowsAffected: false,
					LastInsertID: false,
					Query:        true,
					QueryParams:  true,
				},
			),
		)
		if err != nil {
			return fmt.Errorf("failed to register tracing driver: %v", err)
		}
		ocsql.RegisterAllViews()
	}

	sdb, err := sql.Open(driverName, config.Postgres)
	if err != nil {
		return err
	}

	db := sqlx.NewDb(sdb, driverName)

	if config.MaxIdleConns > 0 {
		db.SetMaxIdleConns(config.MaxIdleConns)
	} else {
		db.SetMaxIdleConns(100)
	}
	if config.MaxOpenConns > 0 {
		db.SetMaxOpenConns(config.MaxOpenConns)
	} else {
		db.SetMaxOpenConns(100)
	}

	instance = &Storage{
		db:  db,
		pqs: make(map[string]*sqlx.Stmt),
	}
	return nil
}

var mu sync.RWMutex

func Gist() *Storage {
	if instance == nil {
		mu.Lock()
		if instance == nil {
			if err := InitConnect(nil); err != nil {
				return &Storage{}
			}
		}
		mu.Unlock()
	}
	return instance
}

func (st *Storage) DB() *sqlx.DB {
	if instance == nil {
		mu.Lock()
		if instance == nil {
			if err := InitConnect(nil); err != nil {
				return &sqlx.DB{}
			}
		}
		mu.Unlock()
	}
	return instance.db
}

func (st *Storage) Query(ctx context.Context) storage.QueryEr {
	if instance == nil {
		st.mu.Lock()
		if instance == nil {
			if err := InitConnect(nil); err != nil {
				return &query{
					ctx: ctx,
				}
			}
		}
		st.mu.Unlock()
	}
	return &query{
		ctx: ctx,
		st:  instance,
	}
}

func (st *Storage) QueryTx(ctx context.Context, f func(qu storage.QueryTxEr) error) (err error) {
	if instance == nil {
		st.mu.Lock()
		if instance == nil {
			if err = InitConnect(nil); err != nil {
				return
			}
		}
		st.mu.Unlock()
	}

	tx, err := instance.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	pgQuery := &queryTx{
		ctx: ctx,
		tx:  tx,
		st:  instance,
	}

	commit := false
	defer func() {
		if r := recover(); r != nil || !commit {
			if r != nil {
				err = fmt.Errorf("transaction panic: %s\n%s", r, string(debug.Stack()))
				_ = pgQuery.Rollback()
			} else if e := pgQuery.Rollback(); e != nil {
				err = e
			}
		} else if commit {
			if e := pgQuery.Commit(); e != nil {
				err = e
			}
		}
	}()

	if err := f(pgQuery); err != nil {
		return err
	}

	commit = true
	return nil
}

// ////

func (st *Storage) querySlice(ctx context.Context, stmt *sqlx.Stmt, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := stmt.Unsafe().QueryxContext(ctx, args...)
	if err != nil {
		return nil, err
	}

	results := make([]map[string]interface{}, 0, 100) // TODO с потолка для оптимизации
	var result map[string]interface{}
	for rows.Next() {
		result = make(map[string]interface{})
		if err := rows.MapScan(result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

func (st *Storage) queryMap(ctx context.Context, stmt *sqlx.Stmt, args ...interface{}) (map[int64]map[string]interface{}, error) {
	rows, err := stmt.Unsafe().QueryxContext(ctx, args...)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, nil
	}
	result := make(map[string]interface{})
	if err := rows.MapScan(result); err != nil {
		return nil, err
	}
	if _, ok := result["id"]; !ok {
		return nil, errors.New("ID is not result exists")
	}

	results := make(map[int64]map[string]interface{})
	results[result["id"].(int64)] = result

	for rows.Next() {
		result = make(map[string]interface{})
		if err := rows.MapScan(result); err != nil {
			return nil, err
		}
		results[result["id"].(int64)] = result
	}
	return results, nil
}
