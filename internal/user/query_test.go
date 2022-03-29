package user

import (
	"context"
	"testing"

	"sample/lib/conf"
	"sample/lib/storage/stpg"
)

func TestQuery(t *testing.T) {
	initTest(t)
	st := stpg.Gist()

	for _, q := range GetQueries() {
		if _, _, err := st.Query(context.Background()).PrepareQuery(q); err != nil {
			t.Log(q)
			t.Fatal(err)
		}
	}
}

func initTest(t *testing.T) {
	var cfg = struct {
		Postgresql stpg.Config `yaml:"psql"`
	}{}
	if err := conf.Get(&cfg, conf.FileConfig, ""); err != nil {
		t.Fatal(err)
	}
	if err := stpg.InitConnect(&cfg.Postgresql); err != nil {
		t.Fatal(err)
	}
}
