package stpg

import (
	"context"
	"testing"
)

// language=sql
const (
	SQL_USER_IN    = "SELECT * FROM users WHERE login = $1 OR login IN ($2) OR login = $3"
	SQL_USER_LIMIT = `
	SELECT
		id, created_at, login, description
	FROM public.users
	LIMIT 100
	`
	SQL_USER = `
	SELECT
		id, created_at, login, description
	FROM public.users
	WHERE
		login = $1
	`
	SQL_USER_INSERT = `INSERT INTO public.users (login, description) VALUES ($1, $2) RETURNING id`
	SQL_USER_UPDATE = `UPDATE public.users SET login = $1, description = $2 WHERE id = $3`
	SQL_USER_UPSERT = `
	INSERT INTO public.users
		(id, login, description)
	VALUES
		($1, $2, $3)
	ON CONFLICT (id) DO UPDATE SET
		login = $2
	`
)

var pgQueries = []string{
	SQL_USER_IN,
	SQL_USER_LIMIT,
	SQL_USER,
	SQL_USER_INSERT,
	SQL_USER_UPDATE,
	SQL_USER_UPSERT,
}

func TestQuery(t *testing.T) {
	var cfg = struct {
		Postgresql Config `yaml:"psql"`
	}{
		Postgresql: getConfig(),
	}
	if err := InitConnect(&cfg.Postgresql); err != nil {
		t.Fatal(err)
	}
	st := Gist()

	for i := range pgQueries {
		if _, _, err := st.Query(context.Background()).PrepareQuery(pgQueries[i]); err != nil {
			t.Log(pgQueries[i])
			t.Fatal(err)
		}
	}
}
