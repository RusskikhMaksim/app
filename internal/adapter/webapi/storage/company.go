package storage

import (
	"app/internal/domain/repository"
	"bytes"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"log"
	"sync"
)

type Bucket struct {
	Name     string
	IsPublic bool
}

type StorageConfig struct {
	Host     string
	Scheme   string
	User     string
	Password string
	Bucket   *Bucket
}

type CompanyRepository struct {
	config *StorageConfig
	client *minio.Client
	mux    *sync.Mutex
}

func NewCompanyRepository(
	config *StorageConfig,
	client *minio.Client,
	mux *sync.Mutex,
) *CompanyRepository {
	return &CompanyRepository{
		config: config,
		client: client,
		mux:    mux,
	}
}

func (r *CompanyRepository) Create(
	ctx context.Context,
	data []byte,
	companyName string,
) (*repository.CompanyRepositoryCreateResponse, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	reader := bytes.NewReader(data)
	name := fmt.Sprintf("%s.json", companyName)

	info, err := r.client.PutObject(
		ctx,
		r.config.Bucket.Name,
		name,
		reader,
		reader.Size(),
		minio.PutObjectOptions{ContentType: "application/json"},
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create company: %s", err.Error())
	}

	log.Printf("%+v", info)

	return &repository.CompanyRepositoryCreateResponse{
		Bucket: r.config.Bucket.Name,
		Name:   name,
	}, nil
}
