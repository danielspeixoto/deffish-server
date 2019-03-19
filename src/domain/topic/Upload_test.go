package topic

import (
	"deffish-server/src/aggregates"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"testing"
)

var q = aggregates.Topic{
	Title: aggregates.Title{
		Value: "A",
	},
	Content: []aggregates.Text{
		{"A"}, {"B"},
	},
}

func TestUploadQuestionSuccessful(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockIRepository(ctrl)
	presenter := NewMockIUploadPresenter(ctrl)

	useCase := Upload{
		Repo: repo,
		Presenter:presenter,
	}

	repo.EXPECT().
		Insert(gomock.Eq(q))
	presenter.EXPECT().
		OnUploaded()

	useCase.Upload(q)
}

func TestUploadQuestionError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockIRepository(ctrl)
	presenter := NewMockIUploadPresenter(ctrl)

	useCase := Upload{
		Repo: repo,
		Presenter:presenter,
	}

	repo.EXPECT().
		Insert(gomock.Eq(q)).
		Return(
			aggregates.Id{},
			errors.New(""))
	presenter.EXPECT().OnError(gomock.Any())

	useCase.Upload(q)
}