package topic

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/topic"
	"errors"
	"log"
)

type Upload struct {
	Repo topic.IRepository
	Presenter topic.IUploadPresenter
}
var _ topic.IUploadUseCase = (*Upload)(nil)

func (useCase Upload) Upload(question aggregates.Topic) {
	id, err := useCase.Repo.Insert(question)
	if err != nil {
		log.Print(err)
		useCase.Presenter.OnError(
			errors.New("could not insert"))
	} else {
		useCase.Presenter.OnUploaded(id)
	}
}
