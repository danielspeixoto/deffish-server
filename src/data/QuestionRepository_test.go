package data

import (
	"bytes"
	"deffish-server/src/domain"
	"io/ioutil"
	"reflect"
	"testing"
)

var repo = NewMongoQuestionRepository(
	"mongodb://localhost:27017",
	"deffishtest",
	"questions")

func TestMain(m *testing.M) {
	err := repo.Drop()
	if err != nil { panic(err) }
	m.Run()
}

func TestInsertedItemsCanBeRetrieved(t *testing.T) {
	pdfBytes, err := ioutil.ReadFile("res/question.pdf")
	if err != nil { t.Fatal(err) }
	question := domain.Question{
		PDF: domain.PDF{
			Content: pdfBytes,
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
	_, err = repo.Insert(question)
	if err != nil { t.Fatal(err) }

	questions, err := repo.Find()
	if err != nil { t.Fatal(err) }

	question.Id = questions[0].Id
	if !reflect.DeepEqual(questions[0], question) {
		t.Fatal("Objects are different")
	}

	mongoPdfBytes := questions[0].PDF.Content
	if bytes.Compare(mongoPdfBytes, pdfBytes) != 0 {
		t.Fatal("PDF is different")
	}


}
