package storage

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Face interface {
	DB() *sqlx.DB
	Query(ctx context.Context) QueryEr
	QueryTx(ctx context.Context, f func(qu QueryTxEr) error) (err error)
}

type QueryEr interface {
	// insert
	ExecInsert(query string, arg ...interface{}) (lastInsertId int64, err error)
	// update, delete, upsert
	Exec(query string, arg ...interface{}) (rowsAffected int64, err error)
	// get one object
	Get(dest interface{}, query string, arg ...interface{}) error
	// get more objects
	Select(dest interface{}, query string, arg ...interface{}) error
	// TODO использовать с осторожностью (требует оптимизации и улучшения)
	QueryMap(query string, arg ...interface{}) (map[int64]map[string]interface{}, error)
	// TODO использовать с осторожностью (требует оптимизации и улучшения)
	QuerySlice(query string, arg ...interface{}) ([]map[string]interface{}, error)
	//
	PrepareQuery(query string, arg ...interface{}) (stmt *sqlx.Stmt, args []interface{}, err error)
}

type QueryTxEr interface {
	QueryEr
	Commit() error
	Rollback() error
}
