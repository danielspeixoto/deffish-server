package data

import (
	"deffish-server/src/domain"
	"reflect"
	"testing"
)

var repo = NewMongoQuestionRepository("mongo://localhost:20171", "deffishtest")

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
			content: []byte{
				1,0,1,0,1,
			},
		},
		Answer: 0,
		Choices: [] domain.Choice{
			{"A"}, {"B"}, {"C"},
		},
		Tags: [] domain.Tag{
			{"matematica"},
			{"enem"},
		},
		Edition: domain.Edition{
			Number: 2017,
		},
	}
	err := repo.Insert(question)
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
