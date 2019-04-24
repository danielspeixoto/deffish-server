package related_video

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/relatedVideo"
)

type UploadUseCase struct {
	Repo relatedVideo.IRepository
}

func (useCase UploadUseCase) FilterByQuestion(id aggregates.Id, start int, count int) ([]aggregates.RelatedVideo, error) {
	if count > 10 {
		count = 10
	}
	return useCase.Repo.FilterByQuestion(id, start, count)
}

var _ relatedVideo.IFilterByQuestionUseCase = (*UploadUseCase)(nil)
