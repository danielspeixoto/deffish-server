package tag

import (
	"deffish-server/src/aggregates"
)

type IRepository interface {
	Insert(aggregates.Tag) (aggregates.Id, error)
	GetByName(string) (aggregates.Tag, error)
	SuggestionsBySubStr(string) ([]aggregates.Tag, error)
}


type IByNameUseCase interface {
	ByName(name string) (aggregates.Tag, error)
}


type IUploadUseCase interface {
	Upload(aggregates.Tag) (aggregates.Id, error)
}

type ISuggestionsBySubStr interface {
	GetSuggestions(string) ([]aggregates.Tag, error)
}