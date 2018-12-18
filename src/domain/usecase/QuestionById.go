package usecase

import (
	"deffish-server/src/domain"
	"deffish-server/src/domain/boundary"
	"deffish-server/src/domain/gateway"
	"github.com/pkg/errors"
)

type QuestionByIdUseCase struct {
	Repo gateway.IQuestionRepository
	Presenter boundary.IQuestionByIdPresenter
}

func (useCase QuestionByIdUseCase) Id(id domain.Id) {
	question, err := useCase.Repo.Id(id)
	if err != nil {
		useCase.Presenter.OnError(errors.New("Could not find"))
	} else {
		useCase.Presenter.OnQuestionReceived(question)
	}
}

var _ boundary.IQuestionByIdUseCase = (*QuestionByIdUseCase)(nil)

