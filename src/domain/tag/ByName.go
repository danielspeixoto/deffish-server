package tag

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/tag"
)

type ByNameUseCase struct {
	Repo tag.IRepository
}

func (useCase ByNameUseCase) ByName(name string) (aggregates.Tag, error) {
	return useCase.Repo.GetByName(name)
}

var _ tag.IByNameUseCase = (*ByNameUseCase)(nil)