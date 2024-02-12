package app

import (
	"app/internal/adapter/webapi/storage"
	"app/internal/interface/server/app"
	"os"
)

func InitConfig(serviceName string) (*Config, error) {
	config := Config{
		AppServer: &app.Config{
			Port: os.Getenv("APPLICATION_PORT"),
		},
		S3Api: &storage.StorageConfig{
			Host:     os.Getenv("S3_API_HOST"),
			User:     os.Getenv("S3_API_USER"),
			Password: os.Getenv("S3_API_PASSWORD"),
			Bucket: &storage.Bucket{
				Name:     os.Getenv("S3_API_BUCKET_NAME"),
				IsPublic: true,
			},
		},
		ServiceName: serviceName,
	}

	return &config, nil
}

type Config struct {
	AppServer   *app.Config
	S3Api       *storage.StorageConfig
	ServiceName string
}
