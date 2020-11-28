package stpg

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type QueryTx struct {
	tx  *sqlx.Tx
	ctx context.Context
	st  *Storage
}

func (qu *QueryTx) Execute(query string, arg []interface{}) error {
	return qu.Exec(query, arg...)
}

func (qu *QueryTx) Exec(query string, arg ...interface{}) error {
	stmt, args, err := qu.PrepareQuery(query, arg...)
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(qu.ctx, args...)
	return err
}

func (qu *QueryTx) Select(dest interface{}, query string, arg ...interface{}) error {
	stmt, args, err := qu.PrepareQuery(query, arg...)
	if err != nil {
		return err
	}
	return stmt.Unsafe().SelectContext(qu.ctx, dest, args...)
}

func (qu *QueryTx) Get(dest interface{}, query string, arg ...interface{}) error {
	stmt, args, err := qu.PrepareQuery(query, arg...)
	if err != nil {
		return err
	}
	return stmt.Unsafe().GetContext(qu.ctx, dest, args...)
}

// TODO реализовать паттерн walker
func (qu *QueryTx) QuerySlice(query string, arg ...interface{}) ([]map[string]interface{}, error) {
	stmt, args, err := qu.PrepareQuery(query, arg...)
	if err != nil {
		return nil, err
	}

	return qu.st.querySlice(qu.ctx, stmt, args...)
}

// TODO реализовать паттерн walker
func (qu *QueryTx) QueryMap(query string, arg ...interface{}) (map[string]map[string]interface{}, error) {
	stmt, args, err := qu.PrepareQuery(query, arg...)
	if err != nil {
		return nil, err
	}
	return qu.st.queryMap(qu.ctx, stmt, args...)
}

func (qu *QueryTx) PrepareQuery(query string, arg ...interface{}) (stmt *sqlx.Stmt, args []interface{}, err error) {
	if len(arg) > 1 || arg[0] != nil {
		if query, args, err = sqlIn(query, arg...); err != nil {
			return
		}
	}
	if stmt, err = qu.prepare(query); err != nil {
		return
	}
	return qu.tx.StmtxContext(qu.ctx, stmt), args, nil
}

func (qu *QueryTx) prepare(query string) (*sqlx.Stmt, error) {
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

func (qu *QueryTx) Commit() error {
	return qu.tx.Commit()
}

func (qu *QueryTx) Rollback() error {
	return qu.tx.Rollback()
}
