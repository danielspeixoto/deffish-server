package question

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/question"
)

type Upload struct {
	Repo question.IRepository
}
var _ question.IUploadUseCase = (*Upload)(nil)

func (useCase Upload) Upload(question aggregates.Question) (aggregates.Id, error) {
	return useCase.Repo.Insert(question)
}