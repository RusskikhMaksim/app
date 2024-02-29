package app

import (
	"app/internal/adapter/webapi/storage"
	"app/internal/consumers"
	"app/internal/domain/service"
	"app/internal/domain/service/producersservices"
	"app/internal/producers"
	"app/internal/registry"
	"app/internal/usecase"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"sync"
)

func NewContainer(
	config *Config,
) (*registry.Container, error) {
	queueClients := newQueueClients(config)
	httpClients, err := newHttpClients(config)
	if err != nil {
		return nil, err
	}

	repositories := newRepositories(config, httpClients)
	services := newServices(repositories, queueClients)
	usecases := newUsecases(services)

	return &registry.Container{
		QueueClientx: queueClients,
		HttpClients:  httpClients,
		Repositories: repositories,
		Services:     services,
		Usecases:     usecases,
	}, nil
}

func newQueueClients(config *Config) *registry.QueueClients {
	return &registry.QueueClients{
		Kafka: &registry.KafkaClients{
			Producers: &registry.Producers{
				FileProcessed: producers.NewFileProcessed(config.Queue.Kafka.ClientConfig),
			},
			Consumers: &registry.Consumers{
				FileProcessed: consumers.NewFileProcessorConsumerGroupHandler(),
			},
		},
	}
}

func newHttpClients(config *Config) (*registry.HttpClients, error) {
	_ = 1
	client, err := minio.New(
		config.S3Api.Host,
		&minio.Options{
			Creds:  credentials.NewStaticV4(config.S3Api.User, config.S3Api.Password, ""),
			Secure: false,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("s3 client creating error: %s", err)
	}

	return &registry.HttpClients{
		StorageClient: client,
	}, nil
}

func newRepositories(
	config *Config,
	clients *registry.HttpClients,
) *registry.Repositories {
	cr := storage.NewCompanyRepository(
		config.S3Api,
		clients.StorageClient,
		&sync.Mutex{},
	)

	return &registry.Repositories{
		CompanyRepository: cr,
	}
}

func newServices(r *registry.Repositories, q *registry.QueueClients) *registry.Services {
	return &registry.Services{
		CompanyService:      service.NewCompanyService(r.CompanyRepository),
		FileExporterService: producersservices.NewFileExporterService(q.Kafka.Producers.FileProcessed),
	}
}

func newUsecases(s *registry.Services) *registry.Usecases {
	return &registry.Usecases{
		Import: usecase.NewImportUsecase(s.CompanyService, s.FileExporterService),
	}
}
