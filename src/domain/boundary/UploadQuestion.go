package boundary

import (
	"deffish-server/src/domain"
)

type IUploadQuestionUseCase interface {
	Upload(domain.Question)
}

type IUploadQuestionPresenter interface {
	OnQuestionUploaded()
	OnError(error)
}

