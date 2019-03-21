package essay

import (
	"deffish-server/src/aggregates"
)

type Comment struct {
	Text string `json:"text"`
}


func fromRequestToComment(comment Comment) aggregates.Comment {
	return aggregates.Comment{
		aggregates.Text{
			comment.Text,
		},
	}
}

func fromCommentToJson(comment aggregates.Comment) Comment {
	return Comment{
		Text:comment.Text.Value,
	}
}

