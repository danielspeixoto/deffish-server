package domain

import (
	"deffish-server/src/aggregates"
	"errors"
	"log"
)

type Upload struct {
	Repo IRepository
	Presenter IUploadPresenter
}
var _ IUploadUseCase = (*Upload)(nil)

func (useCase Upload) Upload(question aggregates.Question) {
	_, err := useCase.Repo.Insert(question)
	if err != nil {
		log.Print(err)
		useCase.Presenter.OnError(
			errors.New("could not insert"))
	} else {
		useCase.Presenter.OnUploaded()
	}
}
