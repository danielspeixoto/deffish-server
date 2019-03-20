package topic

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/topic"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestRandomSetsMaximumValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := topic.NewMockIRepository(ctrl)
	presenter := topic.NewMockIRandomPresenter(ctrl)

	useCase := Random{
		Repo:      repo,
		Presenter: presenter,
		Max:       10,
	}

	repo.EXPECT().
		Random(10).
		Return([]aggregates.Topic{}, nil)

	presenter.EXPECT().
		OnListReceived(gomock.Any())

	useCase.Random(500)
}