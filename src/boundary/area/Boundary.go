package area

import (
	"deffish-server/src/aggregates"
	"github.com/reactivex/rxgo/observable"
)

type IRepository interface {
	Insert(aggregates.Area) (aggregates.Id, error)
	GetByName(string) (aggregates.Area, error)
}

type IByNameUseCase interface {
	ByName(name string) observable.Observable
}


type IUploadUseCase interface {
	Upload(aggregates.Area) observable.Observable
}