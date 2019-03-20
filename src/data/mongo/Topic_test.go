package mongo

import (
	"deffish-server/src/aggregates"
	"reflect"
	"strconv"
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

		t.Run("FindAll FindAll", func(t *testing.T) {
			topics, err := topicRepo.FindAll()
			if err != nil { t.Fatal(err) }

			topic.Id = topics[0].Id
			if !reflect.DeepEqual(topics[0], topic) {
				t.Fatal("Objects are different")
			}
		})
	})
}


func TextTopic_ManyItems(t *testing.T) {
	t.Run("Insert Many Items", func(t *testing.T) {
		for i := 0; i < 5; i++ {
			topic := aggregates.Topic{
				Title: aggregates.Title{
					Value: strconv.Itoa(i),
				},
				Content: []aggregates.Text{
					{"A"}, {"B"},
				},
			}
			_, err := topicRepo.Insert(topic)
			if err != nil { t.Fatal(err) }
		}

		t.Run("Random should give different results each time is run", func(t *testing.T) {
			var randomTopics []aggregates.Topic
			for i := 0; i < 10; i++ {
				questions, err := topicRepo.Random(1)
				if err != nil {
					t.Fatal(err)
				}

				if len(questions) != 1 {
					t.Fail()
				}
				randomTopics = append(randomTopics, questions[0])
			}

			for i := 1; i < len(randomTopics); i++ {
				if randomTopics[i].Title.Value != randomTopics[0].Title.Value {
					break
				}
				if i == len(randomTopics) - 1 {
					t.Errorf("No randomness")
				}
			}
		})

		t.Run("Random without tags should retrieve all", func(t *testing.T) {

			questions, err := topicRepo.Random(100)
			if err != nil { t.Fatal(err) }

			if len(questions) != 5 {
				t.Errorf("Random should return all questions. " +
					"Expected: %v, Got: %v", 5, len(questions))
			}
		})
	})
}







