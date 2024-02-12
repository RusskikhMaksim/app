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
	) error
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
) error {
	err := s.companyRepo.Create(ctx, data, companyName)
	if err != nil {
		return err
	}

	return nil
}
