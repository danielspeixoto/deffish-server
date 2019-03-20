package essay

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/essay"
)

type FilterByTopic struct {
	Repo essay.IRepository
	Presenter essay.IFilterByTopicPresenter
}

func (useCase FilterByTopic) FilterByTopic(topicId aggregates.Id) {
	essays, err := useCase.Repo.FilterByTopic(topicId)
	if err != nil {
		useCase.Presenter.OnError(err)
	} else {
		useCase.Presenter.OnListReceived(essays)
	}
}

var _ essay.IFilterByTopicUseCase = (*FilterByTopic)(nil)