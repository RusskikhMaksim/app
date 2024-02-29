package service

import (
	"app/internal/domain/repository"
	"context"
)

type ICompanyService interface {
	ImportCompany(
		ctx context.Context,
		data []byte,
		companyName string,
	) (*repository.CompanyRepositoryCreateResponse, error)
}

type CompanyService struct {
	companyRepo repository.ICompanyRepository
}

func NewCompanyService(
	companyRepo repository.ICompanyRepository,
) ICompanyService {
	return &CompanyService{companyRepo}
}

func (s *CompanyService) ImportCompany(
	ctx context.Context,
	data []byte,
	companyName string,
) (*repository.CompanyRepositoryCreateResponse, error) {
	r, err := s.companyRepo.Create(ctx, data, companyName)

	if err != nil {
		return nil, err
	}

	return r, nil
}
