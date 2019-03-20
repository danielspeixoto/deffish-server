package topic

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/helpers"
)

type Topic struct {
	Id string `json:"id"`
	Title     string `json:"title"`
	Contents []string `json:"contents"`
}

func fromRequestToTopic(topic Topic) aggregates.Topic {
	return aggregates.Topic{
		Id: aggregates.Id{
			topic.Id,
		},
		Title: aggregates.Title{
			topic.Title,
		},
		Content: helpers.StringArrToTextArr(topic.Contents),
	}
}

func fromTopicToJson(topic aggregates.Topic) Topic {
	return Topic{
		Id: topic.Id.Value,
		Title: topic.Title.Value,
		Contents: helpers.TextArrToStringArray(topic.Content),
	}
}


