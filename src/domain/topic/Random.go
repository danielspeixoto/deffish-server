package essay

import (
	"deffish-server/src/boundary/essay"
	"log"
)

type Random struct {
	Repo      essay.IRepository
	Presenter essay.IRandomPresenter
	Max       int
}

var _ essay.IRandomUseCase = (*Random)(nil)

func (useCase Random) Random(amount int) {
	if amount > useCase.Max {
		log.Printf("An user requested %v essays, max is %v", amount, useCase.Max)
		amount = useCase.Max
	}
	essays, err := useCase.Repo.Random(amount)
	if err != nil {
		log.Print(err)
		useCase.Presenter.OnError(err)
	} else {
		useCase.Presenter.OnListReceived(essays)
	}
}
