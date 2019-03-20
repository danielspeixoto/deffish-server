package topic

import (
	"deffish-server/src/boundary/topic"
	"log"
)

type Random struct {
	Repo      topic.IRepository
	Presenter topic.IRandomPresenter
	Max       int
}

var _ topic.IRandomUseCase = (*Random)(nil)

func (useCase Random) Random(amount int) {
	if amount > useCase.Max {
		log.Printf("An user requested %v topics, max is %v", amount, useCase.Max)
		amount = useCase.Max
	}
	topics, err := useCase.Repo.Random(amount)
	if err != nil {
		log.Print(err)
		useCase.Presenter.OnError(err)
	} else {
		useCase.Presenter.OnListReceived(topics)
	}
}
