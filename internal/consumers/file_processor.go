package consumers

import (
	"app/internal"
	"context"
	"fmt"
	"github.com/IBM/sarama"
)

type FileProcessorConsumerGroupHandler struct {
	group sarama.ConsumerGroup
}

func (FileProcessorConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (FileProcessorConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h FileProcessorConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
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

func NewFileProcessorConsumerGroupHandler() *FileProcessorConsumerGroupHandler {
	return &FileProcessorConsumerGroupHandler{}
}

func Consume(
	handler FileProcessorConsumerGroupHandler,
	group sarama.ConsumerGroup,
	ctx context.Context,
) func(
	handler FileProcessorConsumerGroupHandler,
	group sarama.ConsumerGroup,
	ctx context.Context,
) {
	return func(
		handler FileProcessorConsumerGroupHandler,
		group sarama.ConsumerGroup,
		ctx context.Context,
	) {
		// Track errors
		go func() {
			for err := range group.Errors() {
				fmt.Println("ERROR", err)
			}
		}()

		// Iterate over consumer sessions.
		for {
			topics := []string{internal.TopicFileProcessed}

			err := group.Consume(ctx, topics, handler)
			if err != nil {
				panic(err)
			}
		}
	}
}
