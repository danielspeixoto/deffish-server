package essay

import (
	boundary "deffish-server/src/boundary/essay"
	domain "deffish-server/src/domain/essay"
	routing "deffish-server/src/presentation"
	presentation "deffish-server/src/presentation/essay"
)

func GetRouter(repo boundary.IRepository) routing.IRouter {
	return presentation.Router{
		Controller: func(presenter presentation.Presenter) presentation.Controller {
			uploadEssay := domain.Upload{Repo: repo, Presenter: presenter}
			random := domain.Random{Repo:repo, Presenter:presenter, Max:10}
			byId := domain.ById{Repo:repo, Presenter:presenter}

			return presentation.Controller{
				UploadUseCase: uploadEssay,
				RandomUseCase:random,
				GetById: byId,
			}
		},
	}
}
