package storage

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Face interface {
	Gist() *sqlx.DB
	Query(ctx context.Context) QueryEr
	QueryTx(ctx context.Context, f func(qu QueryTxEr) error) (err error)
}

type QueryEr interface {
	Execute(query string, arg []interface{}) error
	Exec(query string, arg ...interface{}) error
	Get(dest interface{}, query string, arg ...interface{}) error
	Select(dest interface{}, query string, arg ...interface{}) error
	QueryMap(query string, arg ...interface{}) (map[string]map[string]interface{}, error)
	QuerySlice(query string, arg ...interface{}) ([]map[string]interface{}, error)
	PrepareQuery(query string, arg ...interface{}) (stmt *sqlx.Stmt, args []interface{}, err error)
}

type QueryTxEr interface {
	QueryEr
	Commit() error
	Rollback() error
}
