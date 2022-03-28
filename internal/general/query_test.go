package general

import (
	"context"
	"testing"

	"sungora/app/config"
	"sungora/lib/storage/stpg"
)

func TestQuery(t *testing.T) {
	cfg, err := config.Init()
	if err != nil {
		t.Fatal(err)
	}
	if err := stpg.InitConnect(&cfg.Postgresql); err != nil {
		t.Fatal(err)
	}
	st := stpg.Gist()

	for _, q := range GetQueries() {
		if _, _, err := st.Query(context.Background()).PrepareQuery(q); err != nil {
			t.Log(q)
			t.Fatal(err)
		}
	}
}
