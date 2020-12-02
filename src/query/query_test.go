package query

import (
	"context"
	"testing"

	"sungora/lib/app"
	"sungora/lib/storage"
	"sungora/lib/storage/stpg"
)

type PGTest struct {
	storage.Face
}

func TestQuery(t *testing.T) {
	var cfg = struct {
		Postgresql stpg.Config `json:"postgresql"`
	}{}
	if err := app.LoadConfig("conf/config.yaml", &cfg); err != nil {
		t.Fatal(err)
	}
	if err := stpg.InitConnect(&cfg.Postgresql); err != nil {
		t.Fatal(err)
	}
	obj := &PGTest{&stpg.Storage{}}

	for _, q := range GetQueries() {
		if _, _, err := obj.Query(context.Background()).PrepareQuery(q, nil); err != nil {
			t.Log(q)
			t.Fatal(err)
		}
	}
}
