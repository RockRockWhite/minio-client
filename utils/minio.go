package utils

import (
	"context"
	"github.com/RockRockWhite/minio-client/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"log"
)

var _client *minio.Client
var _bucket string

func init() {
	endpoint := config.GetString("Minio.Endpoint")
	accessKeyID := config.GetString("Minio.AccessKeyID")
	secretAccessKey := config.GetString("Minio.SecretAccessKey")
	_bucket = config.GetString("Minio.Bucket")

	// Initialize minio _client object.
	if minioClient, err := minio.New(
		endpoint,
		&minio.Options{
			Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
			Secure: false,
		},
	); err != nil {
		log.Printf("%#v\n", minioClient) // minioClient is now set up
	} else {
		_client = minioClient
	}

	if exists, err := _client.BucketExists(context.Background(), _bucket); !exists {
		if err = _client.MakeBucket(
			context.Background(),
			_bucket,
			minio.MakeBucketOptions{Region: "cn-east-1", ObjectLocking: true},
		); err != nil {
			log.Fatalln(err)
		}
	}
}

func PutObject(objectName string, r io.Reader, objectSize int64) error {
	_, err := _client.PutObject(context.Background(), _bucket, objectName, r, objectSize, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	return err
}

func GetObject(objectName string) (io.Reader, error) {
	object, err := _client.GetObject(context.Background(), _bucket, objectName, minio.GetObjectOptions{})
	return object, err
}
