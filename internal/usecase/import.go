package usecase

import (
	"app/internal/domain/service"
	"context"
	"log"
)

type ImportUsecase struct {
	CompanyService service.ICompanyService
}

func NewImportUsecase(c service.ICompanyService) *ImportUsecase {
	return &ImportUsecase{c}
}

func (u *ImportUsecase) ImportCompany(
	ctx context.Context,
	company string,
	body []byte,
) error {
	err := u.CompanyService.ImportCompany(ctx, body, company)
	if err != nil {
		log.Println(err)
	}
	//dec := json.NewDecoder(bytes.NewReader(body))
	//r := &model.Company{}
	//
	//if err := dec.Decode(r); err != nil {
	//	log.Println(err.Error())
	//}
	//
	//log.Printf("%+v", r)
	/*
		1. Положить файл на с3
		2. Получить etag
		3. Проверить в таблице etag
		4. Если уже есть такой, удалить загруженный файл
	*/

	return nil
}
