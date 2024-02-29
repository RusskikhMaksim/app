package registry

import (
	"app/internal/consumers"
	"app/internal/domain/repository"
	"app/internal/domain/service"
	"app/internal/producers"
	"app/internal/usecase"
	"github.com/minio/minio-go/v7"
)

type Container struct {
	Usecases     *Usecases
	Services     *Services
	Repositories *Repositories
	HttpClients  *HttpClients
	QueueClientx *QueueClients
}

type QueueClients struct {
	Kafka *KafkaClients
}

type KafkaClients struct {
	Producers *Producers
	Consumers *Consumers
}

type Producers struct {
	FileProcessed *producers.FileProcessor
}

type Consumers struct {
	FileProcessed *consumers.FileProcessorConsumerGroupHandler
}

type Usecases struct {
	Import *usecase.ImportUsecase
}

type Services struct {
	CompanyService      service.ICompanyService
	FileExporterService usecase.FileExporterInterface
}

type Repositories struct {
	CompanyRepository repository.ICompanyRepository
}

type HttpClients struct {
	StorageClient *minio.Client
}
