package config

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AWSConfig struct {
	Region string
}

func loadConfig() AWSConfig {
	return AWSConfig{
		Region: os.Getenv("AWS_REGION"),
	}
}

func InitAWS() *s3.Client {
	cfgData := loadConfig()

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(cfgData.Region))
	if err != nil {
		log.Fatal("Fail to load AWS configuration")
	}

	s3Client := s3.NewFromConfig(cfg)

	return s3Client
}
