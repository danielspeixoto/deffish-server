package question

import (
	"deffish-server/src/aggregates"
)

type IRepository interface {
	Insert(aggregates.Question) (aggregates.Id, error)
	RandomByTags(amount int, tags []string) ([]aggregates.Question, error)
	Id(id aggregates.Id) (aggregates.Question, error)
	GetRelatedVideos(id aggregates.Id, start int, count int) ([]aggregates.RelatedVideo, error)
	Add(id aggregates.Id, tag string) error
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

type IAddTagUseCase interface {
	Add(id aggregates.Id, tag string) error
}

type IGetRelatedVideos interface {
	GetRelatedVideos(id aggregates.Id, start int, count int) ([]aggregates.RelatedVideo, error)
}
