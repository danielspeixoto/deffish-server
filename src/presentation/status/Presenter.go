package status

import (
	"deffish-server/src/domain/status"
	"deffish-server/src/presentation/data"
	"encoding/json"
	"net/http"
)

type Presenter struct {
	Writer http.ResponseWriter
}

func (presenter Presenter) Status(status string) {
	presenter.Writer.WriteHeader(http.StatusOK)
	err := json.NewEncoder(presenter.Writer).Encode(data.Response{Status: status})
	if err != nil { panic(err) }
}

var _ status.IPresenter = (*Presenter)(nil)