package presentation

import (
	"deffish-server/src/domain"
	"deffish-server/src/domain/boundary"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Controller struct {
	UploadQuestionUseCase boundary.IUploadQuestionUseCase
	StatusUseCase         boundary.IStatusUseCase
}

type Question struct {
	Pdf []byte
	Answer int
	Choices []string
	Tags []string
}

func (ctrl Controller) Upload(request *http.Request) {
	bodyBytes, err := ioutil.ReadAll(request.Body)
	if err != nil { panic(err) }

	var question Question
	err = json.Unmarshal(bodyBytes, &question)
	if err != nil { panic(err) }

	ctrl.UploadQuestionUseCase.Upload(
		fromRequestToQuestion(question))
}

func (ctrl Controller) Status(request *http.Request) {
	ctrl.StatusUseCase.Status()
}

func fromRequestToQuestion(question Question) domain.Question {
	var choices []domain.Choice
	for _, element := range question.Choices {
		choices = append(choices, domain.Choice{
			Content: element,
		})
	}

	var tags []domain.Tag
	for _, element := range question.Tags {
		tags = append(tags, domain.Tag{
			Name: element,
		})
	}

	return domain.Question{
		PDF: domain.PDF{
			Content: question.Pdf,
		},
		Answer: question.Answer,
		Choices: choices,
		Tags: tags,
	}
}