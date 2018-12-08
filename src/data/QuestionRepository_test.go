package data

import (
	"deffish-server/src/domain"
	"fmt"
	"reflect"
	"testing"
)

var repo = NewMongoQuestionRepository(
	"mongodb://localhost:27017",
	"deffishtest",
	"questions")

func TestMain(m *testing.M) {
	err := repo.Drop()
	if err != nil {
		panic(err)
	}
	m.Run()
}

func TestInsertedItemsCanBeRetrieved(t *testing.T) {
	question := domain.Question{
		PDF: domain.PDF{
			Content: []byte{
				1,0,1,0,1,
			},
		},
		Answer: 0,
		Choices: [] domain.Choice{
			{"A"}, {"B"}, {"C"},
		},
		Tags: [] domain.Tag{
			{"matematica"},
			{"enem2017"},
		},
	}
	id, err := repo.Insert(question)
	fmt.Print(id)
	if err != nil {
		t.Fatal(err)
	}

	questions, err := repo.Find()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(questions[0], question) {
		t.Fatal("Objects are different")
	}
}
