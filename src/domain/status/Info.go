package status

type Info struct {
	Presenter IPresenter
}

func (useCase Info) Info() {
	useCase.Presenter.Status("ok")
}

var _ IInfoUseCase = (*Info)(nil)


