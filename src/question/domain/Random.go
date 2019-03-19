package domain

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/question"
	"log"
)

type Random struct {
	Repo question.IRepository
	Presenter question.IRandomPresenter
	MaxQuestions int
}

var _ question.IRandomUseCase = (*Random)(nil)

func (useCase Random) Random(amount int, tags []aggregates.Tag) {
	if amount > useCase.MaxQuestions {
		log.Printf("An user requested %v questions, max is %v", amount, useCase.MaxQuestions)
		amount = useCase.MaxQuestions
	}
	questions, err := useCase.Repo.Random(amount, tags)
	if err != nil {
		log.Print(err)
		useCase.Presenter.OnError(err)
	} else {
		useCase.Presenter.OnListReceived(questions)
	}
}
