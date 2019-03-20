package topic

import (
	"deffish-server/src/aggregates"
	boundary "deffish-server/src/boundary/topic"
	"encoding/json"
	"net/http"
)

type Presenter struct {
	Writer http.ResponseWriter
}

func (presenter Presenter) OnReceived(topic aggregates.Topic) {
	presenter.Writer.WriteHeader(http.StatusOK)

	err := json.NewEncoder(presenter.Writer).Encode(aggregates.Response{Status: "ok",
		Data: fromTopicToJson(topic)})

	if err != nil { panic(err) }
}

func (presenter Presenter) OnListReceived(topics []aggregates.Topic) {
	presenter.Writer.WriteHeader(http.StatusOK)

	jsonTopics := make([]Topic, 0)
	for _, element := range topics {
		jsonTopics = append(jsonTopics, fromTopicToJson(element))
	}

	err := json.NewEncoder(presenter.Writer).Encode(aggregates.Response{Status: "ok", Data: jsonTopics})
	if err != nil { panic(err) }
}

func (presenter Presenter) OnUploaded() {
	presenter.Writer.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(presenter.Writer).Encode(aggregates.Response{Status: "ok"})
	if err != nil { panic(err) }
}

func (presenter Presenter) OnError(error) {
	presenter.Writer.WriteHeader(http.StatusInternalServerError)
	err := json.NewEncoder(presenter.Writer).Encode(aggregates.Response{Status: "error"})
	if err != nil { panic(err) }
}



var _ boundary.IUploadPresenter = (*Presenter)(nil)
var _ boundary.IRandomPresenter = (*Presenter)(nil)
var _ boundary.IByIdPresenter = (*Presenter)(nil)