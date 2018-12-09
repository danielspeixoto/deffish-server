package usecase

import (
	"deffish-server/src/domain"
	"deffish-server/src/domain/boundary"
	"deffish-server/src/domain/gateway"
)

type RandomQuestionUseCase struct {
	Repo gateway.IQuestionRepository
}

var _ boundary.IRandomQuestionUseCase = (*RandomQuestionUseCase)(nil)

func (RandomQuestionUseCase) Random(amount int, tags []domain.Tag) []domain.Question {
	panic("implement me")
}
