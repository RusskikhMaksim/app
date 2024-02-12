package repository

import "context"

type ICompanyRepository interface {
	Create(
		ctx context.Context,
		data []byte,
		companyName string,
	) error
}
