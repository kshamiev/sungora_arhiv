package general

import (
	"context"
	"testing"

	"sungora/lib/app"
	"sungora/lib/storage/stpg"
	"sungora/src/app/config"
)

func TestQuery(t *testing.T) {
	cfg, err := config.Init(app.ConfigFilePath)
	if err != nil {
		t.Fatal(err)
	}
	if err := stpg.InitConnect(&cfg.Postgresql); err != nil {
		t.Fatal(err)
	}
	st := stpg.Gist()

	for _, q := range GetQueries() {
		if _, _, err := st.Query(context.Background()).PrepareQuery(q, nil); err != nil {
			t.Log(q)
			t.Fatal(err)
		}
	}
}
