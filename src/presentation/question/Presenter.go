package question

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/domain/question"
	"deffish-server/src/presentation/data"
	"encoding/json"
	"net/http"
)

type Presenter struct {
	Writer http.ResponseWriter
}

func (presenter Presenter) OnReceived(question aggregates.Question) {
	presenter.Writer.WriteHeader(http.StatusOK)

	err := json.NewEncoder(presenter.Writer).Encode(data.Response{Status: "ok",
		Data: fromQuestionToJson(question)})

	if err != nil { panic(err) }
}

func (presenter Presenter) OnListReceived(questions []aggregates.Question) {
	presenter.Writer.WriteHeader(http.StatusOK)

	jsonQuestions := make([]Question, 0)
	for _, element := range questions {
		jsonQuestions = append(jsonQuestions, fromQuestionToJson(element))
	}

	err := json.NewEncoder(presenter.Writer).Encode(data.Response{Status: "ok", Data: jsonQuestions})
	if err != nil { panic(err) }
}

func (presenter Presenter) OnUploaded() {
	presenter.Writer.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(presenter.Writer).Encode(data.Response{Status: "ok"})
	if err != nil { panic(err) }
}

func (presenter Presenter) OnError(error) {
	presenter.Writer.WriteHeader(http.StatusInternalServerError)
	err := json.NewEncoder(presenter.Writer).Encode(data.Response{Status: "error"})
	if err != nil { panic(err) }
}



var _ question.IUploadPresenter = (*Presenter)(nil)
var _ question.IRandomPresenter = (*Presenter)(nil)
var _ question.IByIdPresenter = (*Presenter)(nil)