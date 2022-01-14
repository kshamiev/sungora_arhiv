package test

import (
	"os"
	"testing"

	"sungora/lib/typ"
	"sungora/src/miniost"
)

func TestMinio(t *testing.T) {
	t.Skip()
	cfg, ctx := GetEnv()

	err := miniost.CreateBucket("popcorn")
	if err != nil {
		t.Fatal(err)
	}

	fp, err := os.Open(cfg.App.DirWork + "/README.md")
	if err != nil {
		t.Fatal(err)
	}
	fi, err := fp.Stat()
	if err != nil {
		t.Fatal(err)
	}
	err = miniost.PutFile(ctx, "popcorn", typ.UUIDNew().String(), fp, fi.Size())
	if err != nil {
		t.Fatal(err)
	}
}
