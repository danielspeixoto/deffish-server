package mongo

import (
	"deffish-server/src/aggregates"
	"reflect"
	"testing"
)


func TestTopicOneItem(t *testing.T) {
	topic := aggregates.Topic{
		Title: aggregates.Title{
			Value: "A",
		},
		Content: []aggregates.Text{
			{"A"}, {"B"},
		},
	}
	t.Run("Insert One Item", func(t *testing.T) {
		id, err := topicRepo.Insert(topic)
		if err != nil { t.Fatal(err) }

		t.Run("Retrieve using Id", func(t *testing.T) {
			result, err := topicRepo.Id(id)
			if err != nil { t.Fatal(err) }

			if result.Title.Value != "A" {
				t.Errorf("Differs")
			}
		})

		t.Run("Find All", func(t *testing.T) {
			topics, err := topicRepo.Find()
			if err != nil { t.Fatal(err) }

			topic.Id = topics[0].Id
			if !reflect.DeepEqual(topics[0], topic) {
				t.Fatal("Objects are different")
			}
		})
	})
}





