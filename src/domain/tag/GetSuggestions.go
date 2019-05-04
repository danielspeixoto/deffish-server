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
	return useCase.GetSuggestions(query)
}

func (useCase SuggestionsUseCase) GetSuggestions(substr string) ([]aggregates.Tag, error) {
	return useCase.TagRepo.SuggestionsBySubStr(strings.ToLower(substr))
}

var _ tag.ISuggestionsBySubStr = (*SuggestionsUseCase)(nil)