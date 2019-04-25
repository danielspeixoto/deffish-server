package tag

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/tag"
	"strings"
)

type UploadUseCase struct {
	Repo tag.IRepository
}

func (useCase UploadUseCase) Upload(doc aggregates.Tag) (aggregates.Id, error) {
	doc.Name = strings.ToLower(doc.Name)
	return useCase.Repo.Insert(doc)
}

var _ tag.IUploadUseCase = (*UploadUseCase)(nil)