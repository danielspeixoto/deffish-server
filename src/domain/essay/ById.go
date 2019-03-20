package essay

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/essay"
	"github.com/pkg/errors"
)

type ById struct {
	Repo essay.IRepository
	Presenter essay.IByIdPresenter
}

func (useCase ById) Id(id aggregates.Id) {
	question, err := useCase.Repo.Id(id)
	if err != nil {
		useCase.Presenter.OnError(errors.New("Could not find"))
	} else {
		useCase.Presenter.OnReceived(question)
	}
}

var _ essay.IByIdUseCase = (*ById)(nil)

