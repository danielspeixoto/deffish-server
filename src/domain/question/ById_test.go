package question

import (
	"deffish-server/src/aggregates"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestQuestionByIdUseCase_Id(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	presenter := NewMockIByIdPresenter(ctrl)
	repo := NewMockIRepository(ctrl)

	useCase := ById{
		repo,
		presenter,
	}

	id := aggregates.Id{Value: "1"}
	question := aggregates.Question{Id: id}

	repo.EXPECT().
		Id(gomock.Eq(id)).
		Return(question, nil)

	presenter.EXPECT().OnReceived(question)

	useCase.Id(id)
}
