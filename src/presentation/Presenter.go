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
	presenter.Writer.Header().Set("Content-Type", "application/json")
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

func (Presenter) OnError(error) {
	panic("implement me")
}

var _ boundary.IUploadQuestionPresenter = (*Presenter)(nil)
var _ boundary.IStatusPresenter = (*Presenter)(nil)