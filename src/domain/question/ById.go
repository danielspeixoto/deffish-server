package question

import (
	"deffish-server/src/aggregates"
	"github.com/pkg/errors"
)

type ById struct {
	Repo IRepository
	Presenter IByIdPresenter
}

func (useCase ById) Id(id aggregates.Id) {
	question, err := useCase.Repo.Id(id)
	if err != nil {
		useCase.Presenter.OnError(errors.New("Could not find"))
	} else {
		useCase.Presenter.OnReceived(question)
	}
}

var _ IByIdUseCase = (*ById)(nil)

