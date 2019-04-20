package exam

import "deffish-server/src/aggregates"

type IRepository interface {
	Insert(exam aggregates.Exam) (aggregates.Id, error)
	GetByName(name string) (aggregates.Exam, error)
}


type IByNameUseCase interface {
	ByName(name string)
}


type IUploadUseCase interface {
	Upload(aggregates.Exam)
}