package producersservices

import (
	"app/internal"
	"app/internal/producers"
	"app/internal/usecase"
	"context"
	"encoding/json"
	"log"
	"math/rand"
)

type FileExporterService struct {
	FileProcessor *producers.FileProcessor
}

func NewFileExporterService(fileProcessor *producers.FileProcessor) usecase.FileExporterInterface {
	return &FileExporterService{fileProcessor}
}

func (fp *FileExporterService) Send(ctx context.Context, bucket, name string) {
	id := rand.Intn(1000)
	d, err := getValueData(bucket, name)
	if err != nil {
		log.Fatalln(err)
	}

	fp.FileProcessor.Send(internal.TopicFileProcessed, id, d)
}

func getValueData(bucket string, name string) ([]byte, error) {
	d := map[string]string{
		"Bucket": bucket,
		"Name":   name,
	}

	j, err := json.Marshal(&d)
	if err != nil {
		return nil, err
	}

	return j, nil
}
