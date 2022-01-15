package stpg

import (
	"context"
	"testing"

	"sungora/lib/app"
)

// language=sql
const (
	SQL_USER = `
	SELECT
		id, created_at, login, description
	FROM public.users
	WHERE
		login = $1
	`
	SQL_USER_INSERT = `INSERT INTO public.users (id, login, description) VALUES ($1, $2, $3)`
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
	SQL_USER,
	SQL_USER_INSERT,
	SQL_USER_UPDATE,
	SQL_USER_UPSERT,
}

func TestQuery(t *testing.T) {
	var cfg = struct {
		Postgresql Config `json:"postgresql"`
	}{}
	if err := app.LoadConfig(app.ConfigFilePath, &cfg); err != nil {
		t.Fatal(err)
	}
	if err := InitConnect(&cfg.Postgresql); err != nil {
		t.Fatal(err)
	}
	st := Gist()

	for i := range pgQueries {
		if _, _, err := st.Query(context.Background()).PrepareQuery(pgQueries[i], nil); err != nil {
			t.Log(pgQueries[i])
			t.Fatal(err)
		}
	}
}
