package essay

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/helpers"
)

type Essay struct {
	Id string `json:"id"`
	Title     string `json:"title"`
	Text     string `json:"text"`
	Topic     string `json:"topicId"`
	Comments []string `json:"comments"`
}

func fromRequestToEssay(essay Essay) aggregates.Essay {
	return aggregates.Essay{
		Id: aggregates.Id{
			essay.Id,
		},
		Title: aggregates.Title{
			essay.Title,
		},
		Text: aggregates.Text{
			essay.Text,
		},
		Topic: aggregates.Id{
			essay.Topic,
		},
		Comments: helpers.StringArrToCommentArr(essay.Comments),
	}
}

func fromEssayToJson(essay aggregates.Essay) Essay {
	return Essay{
		Id: essay.Id.Value,
		Title: essay.Title.Value,
		Text: essay.Text.Value,
		Topic: essay.Topic.Value,
		Comments: helpers.CommentArrToStringArray(essay.Comments),
	}
}


