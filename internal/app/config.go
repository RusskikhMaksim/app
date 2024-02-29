package app

import (
	"app/internal/adapter/webapi/storage"
	"app/internal/interface/server/app"
	"github.com/IBM/sarama"
	"os"
)

func InitConfig(serviceName string) (*Config, error) {
	var brokers []string
	brokers = append(brokers, os.Getenv("KAFKA_BROKER_FIRST"))
	//brokers = append(brokers, os.Getenv("KAFKA_BROKER_SECOND"))
	//brokers = append(brokers, os.Getenv("KAFKA_BROKER_THIRD"))

	config := Config{
		AppServer: &app.Config{
			Port: os.Getenv("APPLICATION_PORT"),
		},
		Queue: &QueueConfig{
			Kafka: &KafkaConfig{
				ClientConfig:      sarama.NewConfig(),
				BrokersList:       brokers,
				ConsumerGroupName: os.Getenv("KAFKA_CONSUMER_GROUP_NAME"),
			},
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
	Queue       *QueueConfig
	S3Api       *storage.StorageConfig
	ServiceName string
}

type QueueConfig struct {
	Kafka *KafkaConfig
}

type KafkaConfig struct {
	ClientConfig      *sarama.Config
	BrokersList       []string
	ConsumerGroupName string
}
