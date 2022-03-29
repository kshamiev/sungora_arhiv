package minio

import (
	"context"
	"log"
	"os"
	"testing"

	"sample/lib/errs"

	"github.com/google/uuid"
)

func TestMinio(t *testing.T) {
	t.Skip()
	cfg := Config{
		MinioHost:      "localhost:9000",
		MinioAccessKey: "admin",
		MinioSecretKey: "xxx-xxx-xxx",
		MinioSSL:       false,
		MinioRegion:    "eu-east-1",
	}
	if err := Init(&cfg); err != nil {
		log.Fatal(errs.New(err))
	}

	backetID := uuid.New().String()
	if err := CreateBucket(backetID); err != nil {
		t.Fatal(err)
	}

	err := os.WriteFile("test_minio.txt", []byte("mino"), 0o600)
	if err != nil {
		t.Fatal(err)
	}
	fp, err := os.Open("test_minio.txt")
	if err != nil {
		t.Fatal(err)
	}
	fi, err := fp.Stat()
	if err != nil {
		t.Fatal(err)
	}

	err = PutFile(context.Background(), backetID, uuid.New().String(), fp, fi.Size())
	if err != nil {
		t.Fatal(err)
	}
}
