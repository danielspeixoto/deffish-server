package relatedVideo

import "deffish-server/src/aggregates"

type IRepository interface {
	GetRelatedVideos(id aggregates.Id, start int, count int) ([]aggregates.RelatedVideo, error)
}

type IFilterByQuestionUseCase interface {
	GetRelatedVideos(id aggregates.Id, start int, count int) ([]aggregates.RelatedVideo, error)
}
