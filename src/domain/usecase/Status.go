package usecase

import (
	"deffish-server/src/domain/boundary"
	"deffish-server/src/domain/gateway"
)

type StatusUseCase struct {
	Repo gateway.IQuestionRepository
	Presenter boundary.IStatusPresenter
}

func (useCase StatusUseCase) Status() {
	useCase.Presenter.Status("ok")
}

var _ boundary.IStatusUseCase = (*StatusUseCase)(nil)


