package stpg

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type queryTx struct {
	tx  *sqlx.Tx
	ctx context.Context
	st  *Storage
}

func (qu *queryTx) Exec(query string, arg ...interface{}) (int64, error) {
	stmt, args, err := qu.PrepareQuery(query, arg...)
	if err != nil {
		return 0, err
	}
	res, err := stmt.ExecContext(qu.ctx, args...)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	fmt.Println(res.RowsAffected())
	fmt.Println(id, err)
	if id > 0 {
		return id, err
	}
	return res.RowsAffected()
}

func (qu *queryTx) Select(dest interface{}, query string, arg ...interface{}) error {
	stmt, args, err := qu.PrepareQuery(query, arg...)
	if err != nil {
		return err
	}
	return stmt.Unsafe().SelectContext(qu.ctx, dest, args...)
}

func (qu *queryTx) Get(dest interface{}, query string, arg ...interface{}) error {
	stmt, args, err := qu.PrepareQuery(query, arg...)
	if err != nil {
		return err
	}
	return stmt.Unsafe().GetContext(qu.ctx, dest, args...)
}

// TODO реализовать паттерн walker
func (qu *queryTx) QuerySlice(query string, arg ...interface{}) ([]map[string]interface{}, error) {
	stmt, args, err := qu.PrepareQuery(query, arg...)
	if err != nil {
		return nil, err
	}

	return qu.st.querySlice(qu.ctx, stmt, args...)
}

// TODO реализовать паттерн walker
func (qu *queryTx) QueryMap(query string, arg ...interface{}) (map[int64]map[string]interface{}, error) {
	stmt, args, err := qu.PrepareQuery(query, arg...)
	if err != nil {
		return nil, err
	}

	return qu.st.queryMap(qu.ctx, stmt, args...)
}

func (qu *queryTx) PrepareQuery(query string, arg ...interface{}) (stmt *sqlx.Stmt, args []interface{}, err error) {
	if len(arg) > 0 {
		if query, args, err = sqlIn(query, arg...); err != nil {
			return
		}
	}
	if stmt, err = qu.prepare(query); err != nil {
		return
	}
	return qu.tx.StmtxContext(qu.ctx, stmt), args, nil
}

func (qu *queryTx) prepare(query string) (*sqlx.Stmt, error) {
	qu.st.mu.RLock()
	if r, ok := qu.st.pqs[query]; ok {
		qu.st.mu.RUnlock()
		return r, nil
	}
	qu.st.mu.RUnlock()

	qu.st.mu.Lock()
	defer qu.st.mu.Unlock()

	if r, ok := qu.st.pqs[query]; ok {
		return r, nil
	}

	res, err := qu.tx.PreparexContext(qu.ctx, query)
	if err != nil {
		return nil, err
	}

	qu.st.pqs[query] = res
	return res, nil
}

func (qu *queryTx) Commit() error {
	return qu.tx.Commit()
}

func (qu *queryTx) Rollback() error {
	return qu.tx.Rollback()
}
