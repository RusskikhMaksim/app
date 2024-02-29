package producers

import (
	"app/internal"
	"encoding/json"
	"github.com/IBM/sarama"
	"log"
	"math/rand"
	"strconv"
)

var (
	logger  = log.Default()
	brokers = []string{
		"kafka1:9092",
		//"kafka2:9092",
		//"kafka3:9092",
	}
	topic = internal.TopicFileProcessed
)

type FileProcessor struct {
	producer sarama.SyncProducer
}

type entity struct {
	ID   []byte
	Name string
}

func (e *entity) Bytes() []byte {
	bytes, _ := json.Marshal(&e)
	return bytes
}

func newEntity() *entity {
	return &entity{
		ID:   []byte(strconv.Itoa(rand.Intn(100))),
		Name: "test",
	}
}

func (fp *FileProcessor) Send(t string, id int, d []byte) {
	msg := &sarama.ProducerMessage{
		Topic: internal.TopicFileProcessed,
		Key:   sarama.ByteEncoder(strconv.Itoa(id)),
		Value: sarama.ByteEncoder(d),
	}

	_, _, err := fp.producer.SendMessage(msg)
	if err != nil {
		logger.Fatalln(err)
	}

	logger.Printf("sent message '%s' to topic '%s'", id, topic)
}

func NewFileProcessed(cfg *sarama.Config) *FileProcessor {
	cfg.Producer.Partitioner = sarama.NewHashPartitioner
	cfg.Producer.RequiredAcks = sarama.WaitForLocal
	cfg.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, cfg)
	if err != nil {
		logger.Fatalln(err)
	}

	return &FileProcessor{
		producer: producer,
	}
}
