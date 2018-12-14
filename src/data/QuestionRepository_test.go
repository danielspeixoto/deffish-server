package data

import (
	"bytes"
	"deffish-server/src/domain"
	"io/ioutil"
	"reflect"
	"strconv"
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
	err = repo.Drop()
	if err != nil { panic(err) }
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

func TestRandomQuestions(t *testing.T) {
	for i := 0; i < 5; i++ {
		question := domain.Question{
			PDF: domain.PDF{
				Content: []byte{1},
			},
			Answer: i,
			Tags: [] domain.Tag{
				{strconv.Itoa(i)},
			},
		}
		_, err := repo.Insert(question)
		if err != nil { t.Fatal(err) }
	}
	for i := 0; i < 5; i++ {
		question := domain.Question{
			PDF: domain.PDF{
				Content: []byte{1},
			},
			Answer: i + 10,
			Tags: [] domain.Tag{
				{"other"},
				{strconv.Itoa(i + 10)},
			},
		}
		_, err := repo.Insert(question)
		if err != nil { t.Fatal(err) }
	}


	questions, err := repo.Random(100, []domain.Tag{{"other"}})
	if err != nil { t.Fatal(err) }

	if len(questions) != 5 {
		t.Errorf("Random should return questions with tag. " +
			"Expected: %v, Got: %v", 5, len(questions))
	}

	var randomQuestions []domain.Question
	for i := 0; i < 5; i++ {
		questions, err = repo.Random(1, []domain.Tag{{"other"}})
		if err != nil {
			t.Fatal(err)
		}

		if len(questions) != 1 {
			t.Fail()
		}
		randomQuestions = append(randomQuestions, questions[0])
	}

	for i := 1; i < len(randomQuestions); i++ {
		if randomQuestions[i].Answer != randomQuestions[0].Answer {
			break
		}
		if i == len(randomQuestions) - 1 {
			t.Errorf("No randomness")
		}
	}
}

func TestMongoQuestionRepository_RandomNoTags(t *testing.T) {
	for i := 0; i < 5; i++ {
		question := domain.Question{
			PDF: domain.PDF{
				Content: []byte{1},
			},
			Answer: i,
			Tags: [] domain.Tag{
				{strconv.Itoa(i)},
			},
		}
		_, err := repo.Insert(question)
		if err != nil { t.Fatal(err) }
	}

	questions, err := repo.Random(100, []domain.Tag{})
	if err != nil { t.Fatal(err) }

	if len(questions) != 5 {
		t.Errorf("Random should return all questions. " +
			"Expected: %v, Got: %v", 5, len(questions))
	}
}



