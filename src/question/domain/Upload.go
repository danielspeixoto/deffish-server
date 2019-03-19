package domain

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/question"
	"errors"
	"log"
)

type Upload struct {
	Repo question.IRepository
	Presenter question.IUploadPresenter
}
var _ question.IUploadUseCase = (*Upload)(nil)

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
