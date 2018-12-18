package usecase

import (
	"deffish-server/src/domain"
	"deffish-server/src/domain/boundary/mock"
	"deffish-server/src/domain/gateway/mock"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"testing"
)

func TestRandom(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_gateway.NewMockIQuestionRepository(ctrl)
	presenter := mock_boundary.NewMockIRandomQuestionPresenter(ctrl)

	useCase := RandomQuestionUseCase{
		Repo: repo,
		Presenter:presenter,
		MaxQuestions: 10,
	}

	repo.EXPECT().
		Random(2,
			[]domain.Tag{{"tatu"}, {"tuta"},}).
		Return([]domain.Question{
			{
				Id: domain.Id{Value: "1"},
				PDF: domain.PDF{Content: []byte{1}},
				Answer: 0,
			},
			{
				Id: domain.Id{Value: "2"},
				PDF: domain.PDF{Content: []byte{0}},
				Answer: 2,
			},
		}, nil)

	presenter.EXPECT().
		OnQuestionsReceived([]domain.Question{
		{
			Id: domain.Id{Value: "1"},
			PDF: domain.PDF{Content: []byte{1}},
			Answer: 0,
		},
		{
			Id: domain.Id{Value: "2"},
			PDF: domain.PDF{Content: []byte{0}},
			Answer: 2,
		},
	})

	useCase.Random(2, []domain.Tag{{"tatu"}, {"tuta"},})
}

func TestRandomOnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_gateway.NewMockIQuestionRepository(ctrl)
	presenter := mock_boundary.NewMockIRandomQuestionPresenter(ctrl)

	useCase := RandomQuestionUseCase{
		Repo: repo,
		Presenter:presenter,
		MaxQuestions: 10,
	}

	errorReturned := errors.New("error")

	repo.EXPECT().
		Random(5,
			[]domain.Tag{{"tatu"}, {"tuta"},}).
		Return([]domain.Question{}, errorReturned)

	presenter.EXPECT().
		OnError(errorReturned)

	useCase.Random(5, []domain.Tag{{"tatu"}, {"tuta"},})
}

// Parameterized function for minimum value
func TestRandomSetsMaximumValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_gateway.NewMockIQuestionRepository(ctrl)
	presenter := mock_boundary.NewMockIRandomQuestionPresenter(ctrl)

	useCase := RandomQuestionUseCase{
		Repo: repo,
		Presenter:presenter,
		MaxQuestions: 10,
	}

	repo.EXPECT().
		Random(10,
			[]domain.Tag{{"tatu"}, {"tuta"},}).
		Return([]domain.Question{}, nil)

	presenter.EXPECT().
		OnQuestionsReceived(gomock.Any())

	useCase.Random(500, []domain.Tag{{"tatu"}, {"tuta"},})
}