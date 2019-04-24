package question

import (
	"deffish-server/src/aggregates"
)

type IRepository interface {
	Insert(aggregates.Question) (aggregates.Id, error)
	RandomByTags(amount int, tags []string) ([]aggregates.Question, error)
	Id(id aggregates.Id) (aggregates.Question, error)
}

type IRandomByTagsUseCase interface {
	Random(amount int, tags []string) ([]aggregates.Question, error)
}

type IByIdUseCase interface {
	Id(id aggregates.Id) (aggregates.Question, error)
}

type IUploadUseCase interface {
	Upload(aggregates.Question) (aggregates.Id, error)
}