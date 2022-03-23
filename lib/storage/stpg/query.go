package stpg

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type query struct {
	ctx context.Context
	st  *Storage
}

func (qu *query) Exec(query string, arg ...interface{}) (int64, error) {
	stmt, args, err := qu.PrepareQuery(query, arg...)
	if err != nil {
		return 0, err
	}
	res, err := stmt.ExecContext(qu.ctx, args...)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if id > 0 {
		return id, err
	}
	return res.RowsAffected()
}

func (qu *query) Select(dest interface{}, query string, arg ...interface{}) error {
	stmt, args, err := qu.PrepareQuery(query, arg...)
	if err != nil {
		return err
	}
	return stmt.Unsafe().SelectContext(qu.ctx, dest, args...)
}

func (qu *query) Get(dest interface{}, query string, arg ...interface{}) error {
	stmt, args, err := qu.PrepareQuery(query, arg...)
	if err != nil {
		return err
	}
	return stmt.Unsafe().GetContext(qu.ctx, dest, args...)
}

// TODO реализовать паттерн walker
func (qu *query) QuerySlice(query string, arg ...interface{}) ([]map[string]interface{}, error) {
	stmt, args, err := qu.PrepareQuery(query, arg...)
	if err != nil {
		return nil, err
	}

	return qu.st.querySlice(qu.ctx, stmt, args...)
}

// TODO реализовать паттерн walker
func (qu *query) QueryMap(query string, arg ...interface{}) (map[int64]map[string]interface{}, error) {
	stmt, args, err := qu.PrepareQuery(query, arg...)
	if err != nil {
		return nil, err
	}

	return qu.st.queryMap(qu.ctx, stmt, args...)
}

func (qu *query) PrepareQuery(query string, arg ...interface{}) (stmt *sqlx.Stmt, args []interface{}, err error) {
	if len(arg) > 0 {
		if query, args, err = sqlIn(query, arg...); err != nil {
			return
		}
	}
	if stmt, err = qu.prepare(query); err != nil {
		return
	}
	return stmt, args, nil
}

func (qu *query) prepare(query string) (*sqlx.Stmt, error) {
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

	res, err := qu.st.db.PreparexContext(qu.ctx, query)
	if err != nil {
		return nil, err
	}

	qu.st.pqs[query] = res
	return res, nil
}
