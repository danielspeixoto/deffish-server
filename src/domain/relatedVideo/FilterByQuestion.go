package relatedVideo

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/relatedVideo"
)

type FilterByQuestionUseCase struct {
	Repo relatedVideo.IRepository
}

func (useCase FilterByQuestionUseCase) FilterByQuestion(id aggregates.Id, start int, count int) ([]aggregates.RelatedVideo, error) {
	if count > 10 {
		count = 10
	}
	return useCase.Repo.FilterByQuestion(id, start, count)
}

var _ relatedVideo.IFilterByQuestionUseCase = (*FilterByQuestionUseCase)(nil)
