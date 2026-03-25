package s3

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Config 存储S3配置
type Config struct {
	Region     string
	Endpoint   string
	AccessKey  string
	SecretKey  string
	BucketName string
}

// Client 封装S3客户端
type Client struct {
	client     *s3.Client
	bucketName string
}

// NewClient 创建新的S3客户端
func NewS3Client(ctx context.Context, cfg Config) *Client {

	s3cfg := aws.Config{
		Region:       cfg.Region,
		BaseEndpoint: &cfg.Endpoint,
		Credentials: aws.NewCredentialsCache(
			credentials.NewStaticCredentialsProvider(
				cfg.AccessKey,
				cfg.SecretKey,
				"")),
	}

	// build S3 client
	client := s3.NewFromConfig(s3cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	return &Client{
		client:     client,
		bucketName: cfg.BucketName,
	}
}

// PresignUpload 生成预签名上传URL
func (c *Client) PresignUpload(ctx context.Context, key string, expires time.Duration) (string, error) {
	presignClient := s3.NewPresignClient(c.client)
	resp, err := presignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(c.bucketName),
		Key:    aws.String(key),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = expires
	})
	if err != nil {
		return "", err
	}
	return resp.URL, nil
}

// PresignDownload 生成预签名下载URL
func (c *Client) PresignDownload(ctx context.Context, key string, expires time.Duration) (string, error) {
	presignClient := s3.NewPresignClient(c.client)
	resp, err := presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(c.bucketName),
		Key:    aws.String(key),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = expires
	})
	if err != nil {
		return "", err
	}
	return resp.URL, nil
}
