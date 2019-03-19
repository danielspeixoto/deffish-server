package status

type IInfoUseCase interface {
	Info()
}

type IPresenter interface {
	Status(string)
}

