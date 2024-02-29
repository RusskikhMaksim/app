package main

import (
	"app/internal/app"
	"app/internal/consumers"
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/joho/godotenv"
	"log"
)

type exampleConsumerGroupHandler struct{}

func (exampleConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (exampleConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h exampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf(
			"Message topic:%q partition:%d offset:%d\nkey:%s, value:%s",
			msg.Topic,
			msg.Partition,
			msg.Offset,
			msg.Key,
			msg.Value,
		)
		sess.MarkMessage(msg, "")
	}
	return nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
		return
	}
	appConfig, err := app.InitConfig("processor")
	if err != nil {
		log.Fatal("failed to initialize application configuration", err)
	}

	// Start a new consumer group
	group, err := sarama.NewConsumerGroup(
		appConfig.Queue.Kafka.BrokersList,
		appConfig.Queue.Kafka.ConsumerGroupName,
		appConfig.Queue.Kafka.ClientConfig,
	)

	if err != nil {
		panic(err)
	}
	defer func() { _ = group.Close() }()
	//
	//// Track errors
	//go func() {
	//	for err := range group.Errors() {
	//		fmt.Println("ERROR", err)
	//	}
	//}()

	// Iterate over consumer sessions.
	//ctx := context.Background()
	//for {
	//topics := []string{internal.TopicFileProcessed}
	handler := consumers.NewFileProcessorConsumerGroupHandler()
	consumeFunc := consumers.Consume(*handler, group, context.Background())
	consumeFunc(*handler, group, context.Background())
	//err := group.Consume(ctx, topics, handler)
	//if err != nil {
	//	panic(err)
	//}
	//}
}
