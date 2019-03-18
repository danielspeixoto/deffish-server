package domain

import (
	"deffish-server/src/aggregates"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"testing"
)

func TestRandom(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockIRepository(ctrl)
	presenter := NewMockIRandomPresenter(ctrl)

	useCase := Random{
		Repo: repo,
		Presenter:presenter,
		MaxQuestions: 10,
	}

	repo.EXPECT().
		Random(2,
			[]aggregates.Tag{{"tatu"}, {"tuta"},}).
		Return([]aggregates.Question{
			{
				Id: aggregates.Id{Value: "1"},
				PDF: aggregates.PDF{Content: []byte{1}},
				Answer: 0,
			},
			{
				Id: aggregates.Id{Value: "2"},
				PDF: aggregates.PDF{Content: []byte{0}},
				Answer: 2,
			},
		}, nil)

	presenter.EXPECT().
		OnListReceived([]aggregates.Question{
		{
			Id: aggregates.Id{Value: "1"},
			PDF: aggregates.PDF{Content: []byte{1}},
			Answer: 0,
		},
		{
			Id: aggregates.Id{Value: "2"},
			PDF: aggregates.PDF{Content: []byte{0}},
			Answer: 2,
		},
	})

	useCase.Random(2, []aggregates.Tag{{"tatu"}, {"tuta"},})
}

func TestRandomOnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockIRepository(ctrl)
	presenter := NewMockIRandomPresenter(ctrl)

	useCase := Random{
		Repo: repo,
		Presenter:presenter,
		MaxQuestions: 10,
	}

	errorReturned := errors.New("error")

	repo.EXPECT().
		Random(5,
			[]aggregates.Tag{{"tatu"}, {"tuta"},}).
		Return([]aggregates.Question{}, errorReturned)

	presenter.EXPECT().
		OnError(errorReturned)

	useCase.Random(5, []aggregates.Tag{{"tatu"}, {"tuta"},})
}

// Parameterized function for minimum value
func TestRandomSetsMaximumValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockIRepository(ctrl)
	presenter := NewMockIRandomPresenter(ctrl)

	useCase := Random{
		Repo: repo,
		Presenter:presenter,
		MaxQuestions: 10,
	}

	repo.EXPECT().
		Random(10,
			[]aggregates.Tag{{"tatu"}, {"tuta"},}).
		Return([]aggregates.Question{}, nil)

	presenter.EXPECT().
		OnListReceived(gomock.Any())

	useCase.Random(500, []aggregates.Tag{{"tatu"}, {"tuta"},})
}