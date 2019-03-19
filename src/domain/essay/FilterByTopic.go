package essay

import "deffish-server/src/aggregates"

type FilterByTopic struct {
	Repo IRepository
	Presenter IFilterByTopicPresenter
}

func (useCase FilterByTopic) FilterByTopic(topicId aggregates.Id) {
	essays, err := useCase.Repo.FilterByTopic(topicId)
	if err != nil {
		useCase.Presenter.OnError(err)
	} else {
		useCase.Presenter.OnListReceived(essays)
	}
}

var _ IFilterByTopicUseCase = (*FilterByTopic)(nil)