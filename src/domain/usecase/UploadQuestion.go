package usecase

import (
	"deffish-server/src/domain"
	"deffish-server/src/domain/boundary"
	"deffish-server/src/domain/gateway"
	"errors"
)

type UploadQuestionUseCase struct {
	Repo gateway.IQuestionRepository
	Presenter boundary.IUploadQuestionPresenter
}
var _ boundary.IUploadQuestionUseCase = (*UploadQuestionUseCase)(nil)

func (useCase UploadQuestionUseCase) Upload(question domain.Question) {
	_, err := useCase.Repo.Insert(question)
	if err != nil {
		useCase.Presenter.OnError(
			errors.New("could not insert"))
	} else {
		useCase.Presenter.OnQuestionUploaded()
	}
}
