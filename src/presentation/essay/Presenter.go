package essay

import (
	"deffish-server/src/aggregates"
	boundary "deffish-server/src/boundary/essay"
	"deffish-server/src/presentation/data"
	"encoding/json"
	"net/http"
)

type Presenter struct {
	Writer http.ResponseWriter
}

func (presenter Presenter) OnReceived(essay aggregates.Essay) {
	presenter.Writer.WriteHeader(http.StatusOK)

	err := json.NewEncoder(presenter.Writer).Encode(data.Response{Status: "ok",
		Data: fromEssayToJson(essay)})

	if err != nil { panic(err) }
}

func (presenter Presenter) OnListReceived(essays []aggregates.Essay) {
	presenter.Writer.WriteHeader(http.StatusOK)

	jsonEssays := make([]Essay, 0)
	for _, element := range essays {
		jsonEssays = append(jsonEssays, fromEssayToJson(element))
	}

	err := json.NewEncoder(presenter.Writer).Encode(data.Response{Status: "ok", Data: jsonEssays})
	if err != nil { panic(err) }
}

func (presenter Presenter) OnUploaded(id aggregates.Id) {
	presenter.Writer.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(presenter.Writer).Encode(data.Response{Status: "ok", Data: data.Id{id.Value}})
	if err != nil { panic(err) }
}


func (presenter Presenter) OnError(error) {
	presenter.Writer.WriteHeader(http.StatusInternalServerError)
	err := json.NewEncoder(presenter.Writer).Encode(data.Response{Status: "error"})
	if err != nil { panic(err) }
}



var _ boundary.IUploadPresenter = (*Presenter)(nil)
var _ boundary.IRandomPresenter = (*Presenter)(nil)
var _ boundary.IByIdPresenter = (*Presenter)(nil)
var _ boundary.ICommentPresenter = (*Presenter)(nil)
var _ boundary.IFilterByTopicPresenter = (*Presenter)(nil)