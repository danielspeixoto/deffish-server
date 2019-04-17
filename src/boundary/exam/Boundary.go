package essay

import "deffish-server/src/aggregates"

type IRepository interface {
	Insert(essay aggregates.Essay) (aggregates.Id, error)
	Id(id aggregates.Id) (aggregates.Essay, error)
	FilterByTopic(aggregates.Id) ([]aggregates.Essay, error)
	Comment(essayId aggregates.Id, comment aggregates.Comment) error
	Random(amount int) ([]aggregates.Essay, error)
}

type IRandomUseCase interface {
	Random(amount int)
}

type IRandomPresenter interface {
	OnListReceived([]aggregates.Essay)
	OnError(error)
}

type IByIdUseCase interface {
	Id(id aggregates.Id)
}

type IByIdPresenter interface {
	OnReceived(aggregates.Essay)
	OnError(error)
}

type IUploadUseCase interface {
	Upload(aggregates.Essay)
}

type IUploadPresenter interface {
	OnUploaded(id aggregates.Id)
	OnError(error)
}

type IFilterByTopicPresenter interface {
	OnListReceived([]aggregates.Essay)
	OnError(error)
}

type IFilterByTopicUseCase interface {
	FilterByTopic(aggregates.Id)
}

type ICommentPresenter interface {
	OnUploaded(id aggregates.Id)
	OnError(error)
}

type ICommentUseCase interface {
	Comment(essayId aggregates.Id, comment aggregates.Comment)
}
