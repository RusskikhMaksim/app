package repository

import "context"

type CompanyRepositoryCreateResponse struct {
	Bucket string
	Name   string
}

type ICompanyRepository interface {
	Create(
		ctx context.Context,
		data []byte,
		companyName string,
	) (*CompanyRepositoryCreateResponse, error)
}
