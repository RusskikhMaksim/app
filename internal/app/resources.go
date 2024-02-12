package app

import (
	"app/internal/adapter/webapi/storage"
	"app/internal/domain/service"
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
	httpClients, err := newHttpClients(config)
	if err != nil {
		return nil, err
	}

	repositories := newRepositories(config, httpClients)
	services := newServices(repositories)
	usecases := newUsecases(services)

	return &registry.Container{
		HttpClients:  httpClients,
		Repositories: repositories,
		Services:     services,
		Usecases:     usecases,
	}, nil
}

func newHttpClients(config *Config) (*registry.HttpClients, error) {
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

func newServices(r *registry.Repositories) *registry.Services {
	return &registry.Services{
		CompanyService: service.NewCompanyService(r.CompanyRepository),
	}
}

func newUsecases(s *registry.Services) *registry.Usecases {
	return &registry.Usecases{
		Import: usecase.NewImportUsecase(s.CompanyService),
	}
}
