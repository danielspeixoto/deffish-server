package relatedVideo

import "deffish-server/src/aggregates"

type IRepository interface {
	FilterByQuestion(id aggregates.Id, start int, count int) ([]aggregates.RelatedVideo, error)
}

type IFilterByQuestionUseCase interface {
	FilterByQuestion(id aggregates.Id, start int, count int) ([]aggregates.RelatedVideo, error)
}
