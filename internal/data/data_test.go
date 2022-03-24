package data

import (
	"context"
	"log"
	"os"
	"testing"

	"sungora/app/config"
	"sungora/lib/app"
	"sungora/lib/errs"
	"sungora/lib/minio"

	"github.com/google/uuid"
)

func TestMinio(t *testing.T) {
	t.Skip()
	cfg, err := config.Init(app.ConfigFilePath)
	if err != nil {
		t.Fatal(err)
	}

	if err := minio.Init(&cfg.Minio); err != nil {
		log.Fatal(errs.New(err))
	}

	if err := minio.CreateBucket("popcorn"); err != nil {
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
	err = minio.PutFile(context.Background(), "popcorn", uuid.New().String(), fp, fi.Size())
	if err != nil {
		t.Fatal(err)
	}
}
