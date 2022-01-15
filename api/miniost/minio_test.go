package miniost

import (
	"context"
	"log"
	"os"
	"testing"

	"sungora/app/config"
	"sungora/lib/app"
	"sungora/lib/errs"
	"sungora/lib/typ"
)

func TestMinio(t *testing.T) {
	t.Skip()
	cfg, err := config.Init(app.ConfigFilePath)
	if err != nil {
		t.Fatal(err)
	}

	if err := Init(&cfg.Minio); err != nil {
		log.Fatal(errs.NewBadRequest(err))
	}

	if err := CreateBucket("popcorn"); err != nil {
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
	err = PutFile(context.Background(), "popcorn", typ.UUIDNew().String(), fp, fi.Size())
	if err != nil {
		t.Fatal(err)
	}
}
