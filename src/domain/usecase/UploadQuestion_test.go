package usecase

import (
	"deffish-server/src/domain"
	"deffish-server/src/domain/boundary/mock"
	"deffish-server/src/domain/gateway/mock"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"testing"
)

var question = domain.Question{
	PDF: domain.PDF{
		Content: []byte {1, 0},
	},
	Answer: 0,
	Choices: [] domain.Choice{
		{"A"}, {"B"}, {"C"},
	},
	Tags: [] domain.Tag{
		{"matematica"},
		{"enem2017"},
	},
}

func TestUploadQuestionSuccessful(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_gateway.NewMockIQuestionRepository(ctrl)
	presenter := mock_boundary.NewMockIUploadQuestionPresenter(ctrl)

	useCase := UploadQuestionUseCase{
		Repo: repo,
		Presenter:presenter,
	}

	repo.EXPECT().
		Insert(gomock.Eq(question))
	presenter.EXPECT().
		OnQuestionUploaded()

	useCase.Upload(question)
}

func TestUploadQuestionError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_gateway.NewMockIQuestionRepository(ctrl)
	presenter := mock_boundary.NewMockIUploadQuestionPresenter(ctrl)

	useCase := UploadQuestionUseCase{
		Repo: repo,
		Presenter:presenter,
	}

	repo.EXPECT().
		Insert(gomock.Eq(question)).
		Return(
			domain.Id{},
			errors.New(""))
	presenter.EXPECT().OnError(gomock.Any())

	useCase.Upload(question)
}