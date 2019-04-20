package tag

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/tag"
)

type UploadUseCase struct {
	Repo tag.IRepository
}

func (useCase UploadUseCase) Upload(doc aggregates.Tag) (aggregates.Id, error) {
	return useCase.Repo.Insert(doc)
}

var _ tag.IUploadUseCase = (*UploadUseCase)(nil)