package storage

import (
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
) error {
	r.mux.Lock()
	defer r.mux.Unlock()
	reader := bytes.NewReader(data)
	info, err := r.client.PutObject(
		ctx,
		r.config.Bucket.Name,
		fmt.Sprintf("%s.json", companyName),
		reader,
		reader.Size(),
		minio.PutObjectOptions{ContentType: "application/json"},
	)

	if err != nil {
		return fmt.Errorf("failed to create company: %s", err.Error())
	}

	log.Printf("%+v", info)

	return nil
}
