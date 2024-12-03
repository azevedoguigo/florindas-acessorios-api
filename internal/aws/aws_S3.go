package aws

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AwsS3 interface {
	Upload(file multipart.File, filename string) (string, error)
}

type awsS3 struct {
	s3Client *s3.Client
}

func NewAwsS3(s3Client *s3.Client) AwsS3 {
	return awsS3{s3Client: s3Client}
}

func (a awsS3) Upload(file multipart.File, filename string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	timestamp := time.Now().Unix()
	filename = fmt.Sprintf("%d-%s", timestamp, filename)

	_, err := a.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME")),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		return "", err
	}

	fileURL := fmt.Sprintf(
		"https://%s.s3.%s.amazonaws.com/%s",
		os.Getenv("AWS_BUCKET_NAME"),
		os.Getenv("AWS_REGION"),
		filename,
	)

	return fileURL, nil
}
