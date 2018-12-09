package boundary

import (
	"deffish-server/src/domain"
)

type IRandomQuestionUseCase interface {
	Random(amount int, tags []domain.Tag) []domain.Question
}

type IRandomQuestionPresenter interface {
	OnQuestionReceived([]domain.Question)
	OnError(error)
}