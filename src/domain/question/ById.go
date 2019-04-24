package question

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/question"
)

type ById struct {
	Repo question.IRepository
}

func (useCase ById) Id(id aggregates.Id) (aggregates.Question, error){
	return useCase.Repo.Id(id)
}

var _ question.IByIdUseCase = (*ById)(nil)

