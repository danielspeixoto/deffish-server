package domain

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/question"
	"github.com/pkg/errors"
)

type ById struct {
	Repo question.IRepository
	Presenter question.IByIdPresenter
}

func (useCase ById) Id(id aggregates.Id) {
	question, err := useCase.Repo.Id(id)
	if err != nil {
		useCase.Presenter.OnError(errors.New("Could not find"))
	} else {
		useCase.Presenter.OnReceived(question)
	}
}

var _ question.IByIdUseCase = (*ById)(nil)

