package essay

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/essay"
	"errors"
	"log"
)

type Upload struct {
	Repo essay.IRepository
	Presenter essay.IUploadPresenter
}
var _ essay.IUploadUseCase = (*Upload)(nil)

func (useCase Upload) Upload(question aggregates.Essay) {
	id, err := useCase.Repo.Insert(question)
	if err != nil {
		log.Print(err)
		useCase.Presenter.OnError(
			errors.New("could not insert"))
	} else {
		useCase.Presenter.OnUploaded(id)
	}
}
