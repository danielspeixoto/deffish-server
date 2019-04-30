package tag

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/question"
	"deffish-server/src/boundary/tag"
	"strings"
)

type SuggestionsUseCase struct {
	TagRepo tag.IRepository
	QuestionRepo question.IRepository
}

func (useCase SuggestionsUseCase) GetSuggestionsWithQuestions(query string) ([]aggregates.Tag, error) {
	tags, err := useCase.GetSuggestions(query)
	if err != nil {
		return []aggregates.Tag{}, err
	}
	filter := make([]aggregates.Tag, 0)
	for _, t := range tags {
		questions, err := useCase.QuestionRepo.RandomByTags(1, []string { t.Name } )
		if err != nil {
			return tags, err
		}
		if len(questions) > 0 {
			filter = append(filter, t)
		}
	}
	return filter, nil
}

func (useCase SuggestionsUseCase) GetSuggestions(substr string) ([]aggregates.Tag, error) {
	return useCase.TagRepo.SuggestionsBySubStr(strings.ToLower(substr))
}

var _ tag.ISuggestionsBySubStr = (*SuggestionsUseCase)(nil)