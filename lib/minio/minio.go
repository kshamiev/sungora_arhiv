package minio

import (
	"context"
	"io"

	"sample/lib/errs"

	"github.com/minio/minio-go"
)

type Config struct {
	MinioHost      string `yaml:"host"`
	MinioAccessKey string `yaml:"access_key"`
	MinioSecretKey string `yaml:"secret_key"`
	MinioSSL       bool   `yaml:"ssl"`
	MinioRegion    string `yaml:"region"`
}

var conf *Config

var minioClient *minio.Client

func Init(c *Config) (err error) {
	conf = c
	minioClient, err = minio.New(
		c.MinioHost,
		c.MinioAccessKey,
		c.MinioSecretKey,
		c.MinioSSL,
	)
	return err
}

func Gist() *minio.Client {
	return minioClient
}

func CreateBucket(bucket string) error {
	if exists, err := minioClient.BucketExists(bucket); err != nil {
		return errs.New(err, "couldn't check bucket '"+bucket+"'")
	} else if !exists {
		if err := minioClient.MakeBucket(bucket, conf.MinioRegion); err != nil {
			return errs.New(err, "couldn't make bucket '"+bucket+"'")
		}
	}
	return nil
}

func PutFile(ctx context.Context, bucket, objectID string, fileBody io.Reader, size int64) error {
	if err := CreateBucket(bucket); err != nil {
		return errs.New(err)
	}
	opt := minio.PutObjectOptions{}
	if _, err := minioClient.PutObjectWithContext(ctx, bucket, objectID, fileBody, size, opt); err != nil {
		return errs.New(err, "could't put object '"+objectID+"' in bucket '"+bucket+"'")
	}
	return nil
}

func GetFile(ctx context.Context, bucket, objectID string) (io.Reader, error) {
	_, err := minioClient.StatObject(bucket, objectID, minio.StatObjectOptions{})
	if err != nil {
		return nil, errs.New(err, "wrong bucketId or objectId; object "+objectID+" in bucket "+bucket)
	}
	r, err := minioClient.GetObjectWithContext(ctx, bucket, objectID, minio.GetObjectOptions{})
	if err != nil {
		return nil, errs.New(err, "couldn't get object "+objectID+" from bucket "+bucket)
	}
	return r, nil
}

func CopyFile(bucketFrom, objectIDFrom, bucketTo, objectIDTo string) error {
	if _, err := minioClient.StatObject(bucketFrom, objectIDFrom, minio.StatObjectOptions{}); err != nil {
		return errs.New(err, "wrong bucketId or objectId; object "+objectIDFrom+" in bucket "+bucketFrom)
	}
	if _, err := minioClient.StatObject(bucketTo, objectIDTo, minio.StatObjectOptions{}); err == nil {
		return errs.New(nil, "object "+objectIDTo+" in bucket "+bucketTo)
	}
	infoSrc := minio.NewSourceInfo(bucketFrom, objectIDFrom, nil)
	infoDst, err := minio.NewDestinationInfo(bucketTo, objectIDTo, nil, nil)
	if err != nil {
		return errs.New(err, "new dest; object "+objectIDTo+" in bucket "+bucketTo)
	}
	if err := CreateBucket(bucketTo); err != nil {
		return err
	}
	if err := minioClient.CopyObject(infoDst, infoSrc); err != nil {
		return errs.New(err, "couldn't copy object "+objectIDFrom+" from bucket "+bucketFrom)
	}
	return nil
}

func RemoveFile(bucket, objectID string) error {
	if err := minioClient.RemoveObject(bucket, objectID); err != nil {
		return errs.New(err, "couldn't remove object "+objectID+" from bucket "+bucket)
	}
	return nil
}
