package question

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/question"
	"deffish-server/src/boundary/tag"
)

type AddTag struct {
	Repo question.IRepository
	TagRepo tag.IRepository
}

func (useCase AddTag) Add(id aggregates.Id, tag string) error {
	err := useCase.Repo.Add(id, tag)
	if err != nil {
		panic(err)
	}
	return useCase.TagRepo.IncrementCount(tag)
}

var _ question.IAddTagUseCase = (*AddTag)(nil)