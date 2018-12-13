package presentation

import (
	"deffish-server/src/domain/boundary"
	"encoding/json"
	"net/http"
)

type Presenter struct {
	Writer http.ResponseWriter
}

func (presenter Presenter) Status(status string) {
	presenter.Writer.WriteHeader(http.StatusOK)
	err := json.NewEncoder(presenter.Writer).Encode(Response{status})
	if err != nil { panic(err) }
}

type Response struct {
	Status string `json:"status"`
}

func (presenter Presenter) OnQuestionUploaded() {
	presenter.Writer.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(presenter.Writer).Encode(Response{"ok"})
	if err != nil { panic(err) }
}

func (presenter Presenter) OnError(error) {
	presenter.Writer.WriteHeader(http.StatusInternalServerError)
	err := json.NewEncoder(presenter.Writer).Encode(Response{"error"})
	if err != nil { panic(err) }
}

var _ boundary.IUploadQuestionPresenter = (*Presenter)(nil)
var _ boundary.IStatusPresenter = (*Presenter)(nil)