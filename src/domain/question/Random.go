package question

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/question"
	"log"
)

type RandomByTags struct {
	Repo question.IRepository
	MaxQuestions int
}

var _ question.IRandomByTagsUseCase = (*RandomByTags)(nil)

func (useCase RandomByTags) Random(amount int, tags []string) ([]aggregates.Question, error) {
	if amount > useCase.MaxQuestions {
		log.Printf("An user requested %v questions, max is %v", amount, useCase.MaxQuestions)
		amount = useCase.MaxQuestions
	}
	return useCase.Repo.RandomByTags(amount, tags)
}
