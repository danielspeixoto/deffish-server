package topic

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/topic"
	"github.com/pkg/errors"
)

type ById struct {
	Repo topic.IRepository
	Presenter topic.IByIdPresenter
}

func (useCase ById) Id(id aggregates.Id) {
	question, err := useCase.Repo.Id(id)
	if err != nil {
		useCase.Presenter.OnError(errors.New("Could not find"))
	} else {
		useCase.Presenter.OnReceived(question)
	}
}

var _ topic.IByIdUseCase = (*ById)(nil)

