package topic

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/topic"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestQuestionByIdUseCase_Id(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	presenter := topic.NewMockIByIdPresenter(ctrl)
	repo := topic.NewMockIRepository(ctrl)

	useCase := ById{
		repo,
		presenter,
	}

	id := aggregates.Id{Value: "1"}
	topic := aggregates.Topic{Id: id}

	repo.EXPECT().
		Id(gomock.Eq(id)).
		Return(topic, nil)

	presenter.EXPECT().OnReceived(topic)

	useCase.Id(id)
}
