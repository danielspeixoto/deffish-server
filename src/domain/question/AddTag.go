package question

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/question"
)

type AddTag struct {
	Repo question.IRepository
}

func (useCase AddTag) Add(id aggregates.Id, tag string) error {
	return useCase.Repo.Add(id, tag)
}

var _ question.IAddTagUseCase = (*AddTag)(nil)