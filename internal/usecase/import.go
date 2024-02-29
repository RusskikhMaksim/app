package usecase

import (
	"app/internal/domain/service"
	"context"
	"log"
)

type FileExporterInterface interface {
	Send(ctx context.Context, bucket, name string)
}

type ImportUsecase struct {
	CompanyService service.ICompanyService
	FileExporter   FileExporterInterface
}

func NewImportUsecase(
	c service.ICompanyService,
	e FileExporterInterface,
) *ImportUsecase {
	return &ImportUsecase{c, e}
}

func (u *ImportUsecase) ImportCompany(
	ctx context.Context,
	company string,
	body []byte,
) error {
	r, err := u.CompanyService.ImportCompany(ctx, body, company)
	if err != nil {
		log.Println(err)
	}

	u.FileExporter.Send(ctx, r.Bucket, r.Name)
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
