package stpg

import (
	"context"
	"testing"
)

const SQL_USER = `
	SELECT
		id, created_at, login, email
	FROM public.users
	WHERE
		login = $1
	`
const SQL_USER_INSERT = `INSERT INTO public.users (id, login, email) VALUES ($1, $2, $3)`
const SQL_USER_UPDATE = `UPDATE public.users SET login = $1, email = $2 WHERE id = $3`
const SQL_USER_UPSERT = `
	INSERT INTO public.users
		(id, login, email)
	VALUES
		($1, $2, $3)
	ON CONFLICT (id) DO UPDATE SET
		login = $2, email = $3
	`

var pgQueries = []string{
	SQL_USER,
	SQL_USER_INSERT,
	SQL_USER_UPDATE,
	SQL_USER_UPSERT,
}

type PGTest struct {
	Storage
}

func TestQuery(t *testing.T) {
	if err := SetConfig("conf/config.yaml"); err != nil {
		t.Fatal(err)
	}
	obj := &PGTest{}

	for i := range pgQueries {
		if _, _, err := obj.Query(context.Background()).PrepareQuery(pgQueries[i], nil); err != nil {
			t.Log(pgQueries[i])
			t.Fatal(err)
		}
	}
}
