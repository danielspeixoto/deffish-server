package topic

import "deffish-server/src/aggregates"

type IRepository interface {
	Insert(aggregates.Topic) (aggregates.Id, error)
	Id(id aggregates.Id) (aggregates.Topic, error)
}

type IByIdUseCase interface {
	Id(id aggregates.Id)
}

type IByIdPresenter interface {
	OnReceived(aggregates.Topic)
	OnError(error)
}

type IUploadUseCase interface {
	Upload(aggregates.Topic)
}

type IUploadPresenter interface {
	OnUploaded()
	OnError(error)
}