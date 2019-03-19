package question

import (
	"deffish-server/src/aggregates"
)

type IRepository interface {
	Insert(aggregates.Question) (aggregates.Id, error)
	Random(amount int, tags []aggregates.Tag) ([]aggregates.Question, error)
	Id(id aggregates.Id) (aggregates.Question, error)
}

type IRandomUseCase interface {
	Random(amount int, tags []aggregates.Tag)
}

type IRandomPresenter interface {
	OnListReceived([]aggregates.Question)
	OnError(error)
}

type IByIdUseCase interface {
	Id(id aggregates.Id)
}

type IByIdPresenter interface {
	OnReceived(aggregates.Question)
	OnError(error)
}

type IUploadUseCase interface {
	Upload(aggregates.Question)
}

type IUploadPresenter interface {
	OnUploaded()
	OnError(error)
}