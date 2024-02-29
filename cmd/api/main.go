package main

import (
	"app/internal/app"
	"errors"
	"github.com/IBM/sarama"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var BranchName = "development"

var (
	logger  = log.Default()
	brokers = []string{
		"kafka1:9092", // broker_id=1
	}
	topic = "file-processed"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}
	t := os.Getenv("APPLICATION_PORT")
	log.Println("APP PORT", t)

	config, err := app.InitConfig("api")
	if err != nil {
		log.Fatal("failed to initialize application configuration", err)
	}

	createFileProcessedTopic()

	if err := app.Run(config); err != nil {
		log.Fatal("failed to run application", err)
	}
	log.Println("Version: ", BranchName)
}

func createFileProcessedTopic() {
	admin, err := sarama.NewClusterAdmin(brokers, sarama.NewConfig())
	if err != nil {
		logger.Fatalln(err)
	}

	defer func() {
		_ = admin.Close()
	}()

	err = admin.CreateTopic(
		topic,
		&sarama.TopicDetail{
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
		false,
	)
	if err != nil && errors.Unwrap(err) != sarama.ErrTopicAlreadyExists {
		logger.Fatalln(err)
	}
}
