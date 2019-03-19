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

		t.Run("Find All", func(t *testing.T) {
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
	})
}





