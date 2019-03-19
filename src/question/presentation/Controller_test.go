package presentation

import (
	"bytes"
	"deffish-server/src/aggregates"
	"deffish-server/src/question"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)

var q = aggregates.Question{
	PDF: aggregates.PDF{
		Content: []byte {1, 0},
	},
	Answer: 0,
	Choices: [] aggregates.Choice{
		{"A"}, {"B"}, {"C"},
	},
	Tags: [] aggregates.Tag{
		{"matematica"},
		{"enem2017"},
	},
}

func TestController_Upload(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	uploadQuestion := question.NewMockIUploadUseCase(mockCtrl)

	controller := Controller{UploadUseCase: uploadQuestion}

	uploadQuestion.EXPECT().
		Upload(gomock.Eq(q))

	body, err := json.Marshal(map[string]interface{}{
		"pdf":  []byte {1, 0},
		"answer": 0,
		"choices": []string{
			"A", "B", "C",
		},
		"tags": []string{
			"matematica", "enem2017",
		},
	})
	if err != nil { panic(err) }
	request, err := http.NewRequest(
		"POST", "/upload", bytes.NewBuffer(body))
	if err != nil { panic(err) }

	controller.Upload(request)
}

func TestController_Random(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	useCase := question.NewMockIRandomUseCase(mockCtrl)

	controller := Controller{RandomUseCase: useCase}

	useCase.EXPECT().
		Random(2, []aggregates.Tag{{Name: "enem"}, {Name: "matematica"}})

	request, err := http.NewRequest(
		"GET", "/random?amount=2&tags[]=enem&tags[]=matematica", nil)
	if err != nil { panic(err) }

	controller.Random(request)
}

func TestController_RandomEmptyParams(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	useCase := question.NewMockIRandomUseCase(mockCtrl)

	controller := Controller{RandomUseCase: useCase}

	useCase.EXPECT().
		Random(2, gomock.Any())

	request, err := http.NewRequest(
		"GET", "/random", nil)
	if err != nil { panic(err) }

	controller.Random(request)
}