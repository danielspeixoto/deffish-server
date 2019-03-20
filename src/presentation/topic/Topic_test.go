package topic

import (
	"deffish-server/src/aggregates"
	"encoding/json"
)

var example = aggregates.Topic{
	Title: aggregates.Title{
		Value: "A",
	},
	Content: []aggregates.Text{
		{"A"}, {"B"},
	},
}

var exampleJson, _ = json.Marshal(map[string]interface{}{
	"title":  "A",
	"contents": []string{
		"A", "B",
	},
})