package domain

type IUploadQuestionUseCase interface {
	Upload(question Question)
}

type IUploadQuestionPresenter interface {
	OnQuestionUploaded()
	OnError(err error)
}

type IQuestionRepository interface {
	Insert(question Question)
	Random() Question
}

type UploadQuestionUseCase struct {
	repo IQuestionRepository
}

func (useCase UploadQuestionUseCase) Upload(question Question) {

}
