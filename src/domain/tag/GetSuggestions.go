package tag

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/tag"
)

type SuggestionsUseCase struct {
	Repo tag.IRepository
}


func (useCase SuggestionsUseCase) GetSuggestions(substr string) ([]aggregates.Tag, error) {
	return useCase.Repo.SuggestionsBySubStr(substr)
}

var _ tag.ISuggestionsBySubStr = (*SuggestionsUseCase)(nil)