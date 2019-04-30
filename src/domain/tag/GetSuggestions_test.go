package tag

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/question"
	"deffish-server/src/boundary/tag"
	"github.com/golang/mock/gomock"
	"strconv"
	"testing"
)

func TestSuggestionsUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tRepo := tag.NewMockIRepository(ctrl)
	qRepo := question.NewMockIRepository(ctrl)

	useCase := SuggestionsUseCase{
		tRepo, qRepo,
	}

	t.Run("With Question", func(t *testing.T) {
		tRepo.EXPECT().
			SuggestionsBySubStr("ab").
			Return([]aggregates.Tag{
				{
					Id: aggregates.Id{
						"1",
					},
					Name: "abc",
				},
				{
					Id: aggregates.Id{
						"2",
					},
					Name: "dcabc",
				},
			}, nil)
		qRepo.EXPECT().
			RandomByTags(gomock.Any(), []string{"abc"}).
			Return([]aggregates.Question{}, nil)
		qRepo.EXPECT().
			RandomByTags(gomock.Any(), []string{"dcabc"}).
			Return([]aggregates.Question{
				{
					Number: 1,
				},
			}, nil)
		tags, err := useCase.GetSuggestionsWithQuestions("Ab")
		if err != nil {
			panic(err)
		}
		if len(tags) != 1 {
			t.Fatalf("Length: " + strconv.Itoa(len(tags)))
		}
		if tags[0].Name != "dcabc" {
			t.Fatalf("Incorrect result")
		}
	})
}
