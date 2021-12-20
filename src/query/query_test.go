package query

import (
	"context"
	"testing"

	"sungora/lib/app"
	"sungora/lib/storage"
	"sungora/lib/storage/pgsql"
)

type PGTest struct {
	storage.Face
}

func TestQuery(t *testing.T) {
	var cfg = struct {
		Postgresql pgsql.Config `json:"postgresql"`
	}{}
	if err := app.LoadConfig("conf/config.yaml", &cfg); err != nil {
		t.Fatal(err)
	}
	if err := pgsql.InitConnect(&cfg.Postgresql); err != nil {
		t.Fatal(err)
	}
	st := pgsql.Gist()

	for _, q := range GetQueries() {
		if _, _, err := st.Query(context.Background()).PrepareQuery(q, nil); err != nil {
			t.Log(q)
			t.Fatal(err)
		}
	}
}
