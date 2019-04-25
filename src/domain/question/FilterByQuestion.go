package question

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/question"
	"deffish-server/src/boundary/relatedVideo"
)

type FilterByQuestionUseCase struct {
	Repo relatedVideo.IRepository
}

func (useCase FilterByQuestionUseCase) GetRelatedVideos(id aggregates.Id, start int, count int) ([]aggregates.RelatedVideo, error) {
	if count > 10 {
		count = 10
	}
	return useCase.Repo.GetRelatedVideos(id, start, count)
}

var _ question.IGetRelatedVideos = (*FilterByQuestionUseCase)(nil)
