package query

import (
	"context"
	"testing"

	"sungora/lib/storage"
	"sungora/lib/storage/stpg"
)

type PGTest struct {
	storage.Face
}

func TestQuery(t *testing.T) {
	if err := stpg.SetConfig("conf/config.yaml"); err != nil {
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
