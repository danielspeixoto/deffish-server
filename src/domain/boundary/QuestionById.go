package boundary

import "deffish-server/src/domain"

type IQuestionByIdUseCase interface {
	Id(id domain.Id)
}

type IQuestionByIdPresenter interface {
	OnQuestionReceived(domain.Question)
	OnError(error)
}