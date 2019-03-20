package essay

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/essay"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestRandomSetsMaximumValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := essay.NewMockIRepository(ctrl)
	presenter := essay.NewMockIRandomPresenter(ctrl)

	useCase := Random{
		Repo:      repo,
		Presenter: presenter,
		Max:       10,
	}

	repo.EXPECT().
		Random(10).
		Return([]aggregates.Essay{}, nil)

	presenter.EXPECT().
		OnListReceived(gomock.Any())

	useCase.Random(500)
}