package presentation

import (
	"deffish-server/src/domain"
	"deffish-server/src/domain/boundary"
	"encoding/json"
	"net/http"
)

type Presenter struct {
	Writer http.ResponseWriter
}

func (presenter Presenter) OnQuestionReceived(questions []domain.Question) {
	presenter.Writer.WriteHeader(http.StatusOK)

	var jsonQuestions []Question
	for _, element := range questions {
		jsonQuestions = append(jsonQuestions, fromQuestionToJson(element))
	}

	err := json.NewEncoder(presenter.Writer).Encode(Response{Status: "ok", Data: jsonQuestions})
	if err != nil { panic(err) }
}

func (presenter Presenter) Status(status string) {
	presenter.Writer.WriteHeader(http.StatusOK)
	err := json.NewEncoder(presenter.Writer).Encode(Response{Status: status})
	if err != nil { panic(err) }
}


type Response struct {
	Status string `json:"status"`
	Data interface{} `json:"data"`
}

func (presenter Presenter) OnQuestionUploaded() {
	presenter.Writer.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(presenter.Writer).Encode(Response{Status: "ok"})
	if err != nil { panic(err) }
}

func (presenter Presenter) OnError(error) {
	presenter.Writer.WriteHeader(http.StatusInternalServerError)
	err := json.NewEncoder(presenter.Writer).Encode(Response{Status: "error"})
	if err != nil { panic(err) }
}



var _ boundary.IUploadQuestionPresenter = (*Presenter)(nil)
var _ boundary.IRandomQuestionPresenter = (*Presenter)(nil)
var _ boundary.IStatusPresenter = (*Presenter)(nil)