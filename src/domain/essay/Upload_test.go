package essay

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/essay"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"testing"
)

var q = aggregates.Essay{
	Title: aggregates.Title{
		Value: "A",
	},
	Text: aggregates.Text{"abcdef"},
	Topic:aggregates.Id{"1"},
	Comments:[]aggregates.Comment{
		{aggregates.Text{"A"}},
		{aggregates.Text{"B"}},
	},
}

func TestUploadSuccessful(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := essay.NewMockIRepository(ctrl)
	presenter := essay.NewMockIUploadPresenter(ctrl)

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

func TestUploadError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := essay.NewMockIRepository(ctrl)
	presenter := essay.NewMockIUploadPresenter(ctrl)

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