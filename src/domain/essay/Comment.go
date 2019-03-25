package essay

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/essay"
)

type Comment struct {
	Repo essay.IRepository
	Presenter essay.ICommentPresenter
}

func (useCase Comment) Comment(essayId aggregates.Id, comment aggregates.Comment) {
	err := useCase.Repo.Comment(essayId, comment)
	if err != nil {
		useCase.Presenter.OnError(err)
	} else {
		useCase.Presenter.OnUploaded(essayId)
	}
}

var _ essay.ICommentUseCase = (*Comment)(nil)