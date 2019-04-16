package question

import (
	"deffish-server/src/aggregates"
)

type IRepository interface {
	Insert(aggregates.Question) (aggregates.Id, error)
	RandomByDomain(amount int, domain string) ([]aggregates.Question, error)
	RandomByTags(amount int, tags []string) ([]aggregates.Question, error)
	Id(id aggregates.Id) (aggregates.Question, error)
}

type IRandomByDomainUseCase interface {
	Random(amount int, domain string)
}

type IRandomByTagsUseCase interface {
	Random(amount int, tags []string)
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