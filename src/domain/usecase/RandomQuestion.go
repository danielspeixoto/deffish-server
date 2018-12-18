package usecase

import (
	"deffish-server/src/domain"
	"deffish-server/src/domain/boundary"
	"deffish-server/src/domain/gateway"
	"log"
)

type RandomQuestionUseCase struct {
	Repo gateway.IQuestionRepository
	Presenter boundary.IRandomQuestionPresenter
	MaxQuestions int
}

var _ boundary.IRandomQuestionUseCase = (*RandomQuestionUseCase)(nil)

func (useCase RandomQuestionUseCase) Random(amount int, tags []domain.Tag) {
	if amount > useCase.MaxQuestions {
		log.Printf("An user requested %v questions, max is %v", amount, useCase.MaxQuestions)
		amount = useCase.MaxQuestions
	}
	questions, err := useCase.Repo.Random(amount, tags)
	if err != nil {
		log.Print(err)
		useCase.Presenter.OnError(err)
	} else {
		useCase.Presenter.OnQuestionsReceived(questions)
	}
}
