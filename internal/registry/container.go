package registry

import (
	"app/internal/domain/repository"
	"app/internal/domain/service"
	"app/internal/usecase"
	"github.com/minio/minio-go/v7"
)

type Container struct {
	Usecases     *Usecases
	Services     *Services
	Repositories *Repositories
	HttpClients  *HttpClients
}

type Usecases struct {
	Import *usecase.ImportUsecase
}

type Services struct {
	CompanyService service.ICompanyService
}

type Repositories struct {
	CompanyRepository repository.ICompanyRepository
}

type HttpClients struct {
	StorageClient *minio.Client
}
