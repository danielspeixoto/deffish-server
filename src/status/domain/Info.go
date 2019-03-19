package domain

import "deffish-server/src/status"

type Info struct {
	Presenter status.IPresenter
}

func (useCase Info) Info() {
	useCase.Presenter.Status("ok")
}

var _ status.IInfoUseCase = (*Info)(nil)


