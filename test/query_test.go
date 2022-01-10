package test

import (
	"testing"

	"sungora/lib/storage/pgsql"
	"sungora/src/general"
	"sungora/src/user"
)

func TestQuery(t *testing.T) {
	// var cfg = struct {
	// 	Postgresql pgsql.Config `json:"postgresql"`
	// }{}
	// if err := app.LoadConfig("conf/config.yaml", &cfg); err != nil {
	// 	t.Fatal(err)
	// }

	cfg, ctx := GetEnv()

	if err := pgsql.InitConnect(&cfg.Postgresql); err != nil {
		t.Fatal(err)
	}
	st := pgsql.Gist()

	for _, q := range general.GetQueries() {
		if _, _, err := st.Query(ctx).PrepareQuery(q, nil); err != nil {
			t.Log(q)
			t.Fatal(err)
		}
	}
	for _, q := range user.GetQueries() {
		if _, _, err := st.Query(ctx).PrepareQuery(q, nil); err != nil {
			t.Log(q)
			t.Fatal(err)
		}
	}
}