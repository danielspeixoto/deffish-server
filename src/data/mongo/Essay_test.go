package mongo

import (
	"deffish-server/src/aggregates"
	"reflect"
	"strconv"
	"testing"
)

func TestEssayOneItem(t *testing.T) {
	essay := aggregates.Essay{
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
	t.Run("Insert One Item", func(t *testing.T) {
		id, err := essayRepo.Insert(essay)
		if err != nil { t.Fatal(err) }

		t.Run("Retrieve using Id", func(t *testing.T) {
			result, err := essayRepo.Id(id)
			if err != nil { t.Fatal(err) }

			if result.Title.Value != "A" {
				t.Errorf("Title Differs")
			}

			if result.Text.Value != "abcdef" {
				t.Errorf("Text Differs")
			}
		})

		t.Run("FindAll FindAll", func(t *testing.T) {
			essays, err := essayRepo.Find()
			if err != nil { t.Fatal(err) }

			essay.Id = essays[0].Id
			if !reflect.DeepEqual(essays[0], essay) {
				t.Fatal("Objects are different")
			}
		})

		t.Run("Comment", func(t *testing.T) {
			_ = essayRepo.Comment(id, aggregates.Comment{aggregates.Text{"C"}})
			result, _ := essayRepo.Id(id)
			if len(result.Comments) != 3 {
				t.Errorf("Comment arr should be of size 3")
			}
			if result.Comments[2].Text.Value != "C" {
				t.Errorf("Comment inserted not correct")
			}
		})
	})
}

func TextEssay_ManyItems(t *testing.T) {
	t.Run("Insert Many Items", func(t *testing.T) {
		for i := 0; i < 5; i++ {
			essay := aggregates.Essay{
				Title: aggregates.Title{
					Value: "A",
				},
				Text:  aggregates.Text{"abcdef"},
				Topic: aggregates.Id{strconv.Itoa(i)},
				Comments: []aggregates.Comment{
					{aggregates.Text{"A"}},
					{aggregates.Text{"B"}},
				},
			}
			_, err := essayRepo.Insert(essay)
			if err != nil { t.Fatal(err) }
		}
		essay := aggregates.Essay{
			Title: aggregates.Title{
				Value: "B",
			},
			Text:  aggregates.Text{"abcdef"},
			Topic: aggregates.Id{"0"},
			Comments: []aggregates.Comment{
				{aggregates.Text{"A"}},
				{aggregates.Text{"B"}},
			},
		}
		_, err := essayRepo.Insert(essay)
		if err != nil { t.Fatal(err) }

		t.Run("Filter by topic", func(t *testing.T) {
			essays, err := essayRepo.FilterByTopic(aggregates.Id{"0"})
			if err != nil { t.Fatal(err) }
			if len(essays) != 2 {
				t.Fatalf("Size incorrect")
			}
			if essays[1].Title.Value != "B" {
				t.Fatalf("Title incorrect")
			}
		})

		t.Run("Random should give different results each time is run", func(t *testing.T) {
			var randomEssays []aggregates.Essay
			for i := 0; i < 10; i++ {
				questions, err := essayRepo.Random(1)
				if err != nil {
					t.Fatal(err)
				}

				if len(questions) != 1 {
					t.Fail()
				}
				randomEssays = append(randomEssays, questions[0])
			}

			for i := 1; i < len(randomEssays); i++ {
				if randomEssays[i].Topic.Value != randomEssays[0].Topic.Value {
					break
				}
				if i == len(randomEssays) - 1 {
					t.Errorf("No randomness")
				}
			}
		})

		t.Run("Random without tags should retrieve all", func(t *testing.T) {

			questions, err := essayRepo.Random(100)
			if err != nil { t.Fatal(err) }

			if len(questions) != 6 {
				t.Errorf("Random should return all questions. " +
					"Expected: %v, Got: %v", 6, len(questions))
			}
		})
	})
}





