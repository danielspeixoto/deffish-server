package essay

import (
	"deffish-server/src/aggregates"
	"encoding/json"
)

var example = aggregates.Essay{
	Title: aggregates.Title{
		Value: "A",
	},
	Text: aggregates.Text{"abcdef"},
	Topic:aggregates.Id{"1"},
	Comments:[]aggregates.Comment{
		{aggregates.Text{"A"}},
		{aggregates.Text{"B"}},
	},
}

var exampleJson, _ = json.Marshal(map[string]interface{}{
	"title":  "A",
	"text": "abcdef",
	"topicId": "1",
	"comments": []string{
		"A", "B",
	},
})