package presentation

import (
	"bytes"
	"deffish-server/src/domain"
	"deffish-server/src/domain/boundary/mock"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"net/http"
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

func TestController_Upload(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	uploadQuestion := mock_boundary.NewMockIUploadQuestionUseCase(mockCtrl)

	controller := Controller{UploadQuestionUseCase: uploadQuestion}

	uploadQuestion.EXPECT().
		Upload(gomock.Eq(question))

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

	useCase := mock_boundary.NewMockIRandomQuestionUseCase(mockCtrl)

	controller := Controller{RandomQuestionUseCase: useCase}

	useCase.EXPECT().
		Random(2, []domain.Tag{{Name: "enem"}, {Name: "matematica"}})

	request, err := http.NewRequest(
		"GET", "/random?amount=2&tags[]=enem&tags[]=matematica", nil)
	if err != nil { panic(err) }

	controller.Random(request)
}