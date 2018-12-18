package usecase

import (
	"deffish-server/src/domain"
	"deffish-server/src/domain/boundary/mock"
	"deffish-server/src/domain/gateway/mock"
	"github.com/golang/mock/gomock"
	"testing")

func TestQuestionByIdUseCase_Id(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	presenter := mock_boundary.NewMockIQuestionByIdPresenter(ctrl)
	repo := mock_gateway.NewMockIQuestionRepository(ctrl)

	useCase := QuestionByIdUseCase{
		repo,
		presenter,
	}

	id := domain.Id{Value: "1"}
	question := domain.Question{Id: id}

	repo.EXPECT().
		Id(gomock.Eq(id)).
		Return(question, nil)

	presenter.EXPECT().OnQuestionReceived(question)

	useCase.Id(id)
}
