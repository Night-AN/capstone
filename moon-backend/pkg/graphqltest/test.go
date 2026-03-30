package graphqltest

import (
	"context"
	"log"
	"moon/ent"
	"moon/pkg/s3"

	_ "github.com/lib/pq"
)

const (
	RUSTFS_ACCESSKEY = "7EubftlJSGnVZm4C8vKY"
	RUSTFS_SECRETKEY = "SwyYG39BnCPDsMTduvkZeilgEtfjQI20741HKraX"
	RUSTFS_RIGION    = "cn-south-1"
	RUSTFS_ENDPOINT  = "http://localhost:9000"
	RUSTFS_BUCKET    = "capstone"
)

// NewTestEntClient 创建并返回一个测试用的 ent 客户端
func NewTestEntClient() *ent.Client {
	// 创建 ent 客户端
	client, err := ent.Open("postgres", "host=localhost port=5432 user=capstone password=capstone dbname=capstone sslmode=disable")
	if err != nil {
		log.Fatalf("opening connection: %v", err)
	}

	return client
}

// NewTestS3Client 创建并返回一个测试用的 S3 客户端
func NewTestS3Client() *s3.Client {
	// 创建 S3 客户端
	s3Client := s3.NewS3Client(context.Background(), s3.Config{
		Region:     RUSTFS_RIGION,
		Endpoint:   RUSTFS_ENDPOINT,
		AccessKey:  RUSTFS_ACCESSKEY,
		SecretKey:  RUSTFS_SECRETKEY,
		BucketName: RUSTFS_BUCKET,
	})

	return s3Client
}
