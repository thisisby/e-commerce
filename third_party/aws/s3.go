package aws

import (
	"context"
	"fmt"
	config2 "ga_marketplace/internal/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/labstack/gommon/log"
	"mime/multipart"
)

type S3Client struct {
	client *s3.Client
}

func NewS3Client() *S3Client {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(config2.AppConfig.AwsBucketRegion),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(config2.AppConfig.AwsAccessKeyID, config2.AppConfig.AwsAccessKeySecret, "")),
	)
	if err != nil {
		panic(err)
	}

	client := s3.NewFromConfig(cfg)
	return &S3Client{
		client: client,
	}
}

func (s *S3Client) ListBuckets() []types.Bucket {
	result, err := s.client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	var buckets []types.Bucket
	if err != nil {
		log.Printf("failed to list buckets, %v", err)
	} else {
		buckets = result.Buckets
	}

	return buckets
}

func (s *S3Client) CreateBucket(bucketName string) {
	_, err := s.client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Printf("failed to create bucket, %v", err)
	}
}

func (s *S3Client) UploadFile(key string, file *multipart.FileHeader) (string, error) {
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = s.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(config2.AppConfig.AwsBucketName),
		Key:    aws.String(key),
		Body:   f,
		ACL:    "public-read",
	})
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("https://berish91-bucket-public.s3.amazonaws.com/%s", key)

	return url, nil
}
