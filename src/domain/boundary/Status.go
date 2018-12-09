package boundary

type IStatusUseCase interface {
	Status()
}

type IStatusPresenter interface {
	Status(string)
}

