package data

import (
	"bytes"
	"deffish-server/src/aggregates"
	"io/ioutil"
	"reflect"
	"strconv"
	"testing"
)

var repo = NewMongoRepository(
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
	question := aggregates.Question{
		PDF: aggregates.PDF{
			Content: pdfBytes,
		},
		Answer: 0,
		Choices: [] aggregates.Choice{
			{"A"}, {"B"}, {"C"},
		},
		Tags: [] aggregates.Tag{
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
		question := aggregates.Question{
			PDF: aggregates.PDF{
				Content: []byte{1},
			},
			Answer: i,
			Tags: [] aggregates.Tag{
				{strconv.Itoa(i)},
			},
		}
		_, err := repo.Insert(question)
		if err != nil { t.Fatal(err) }
	}
	for i := 0; i < 5; i++ {
		question := aggregates.Question{
			PDF: aggregates.PDF{
				Content: []byte{1},
			},
			Answer: i + 10,
			Tags: [] aggregates.Tag{
				{"other"},
				{strconv.Itoa(i + 10)},
			},
		}
		_, err := repo.Insert(question)
		if err != nil { t.Fatal(err) }
	}


	questions, err := repo.Random(100, []aggregates.Tag{{"other"}})
	if err != nil { t.Fatal(err) }

	if len(questions) != 5 {
		t.Errorf("Random should return questions with tag. " +
			"Expected: %v, Got: %v", 5, len(questions))
	}

	var randomQuestions []aggregates.Question
	for i := 0; i < 5; i++ {
		questions, err = repo.Random(1, []aggregates.Tag{{"other"}})
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
		question := aggregates.Question{
			PDF: aggregates.PDF{
				Content: []byte{1},
			},
			Answer: i,
			Tags: [] aggregates.Tag{
				{strconv.Itoa(i)},
			},
		}
		_, err := repo.Insert(question)
		if err != nil { t.Fatal(err) }
	}

	questions, err := repo.Random(100, []aggregates.Tag{})
	if err != nil { t.Fatal(err) }

	if len(questions) != 5 {
		t.Errorf("Random should return all questions. " +
			"Expected: %v, Got: %v", 5, len(questions))
	}
}

func TestMongoQuestionRepository_CanRetrieveUsingId(t *testing.T) {
	question := aggregates.Question{
		PDF: aggregates.PDF{
			Content: []byte{1},
	},
		Answer: 1,
	}
	id, err := repo.Insert(question)
	if err != nil { t.Fatal(err) }
	result, err := repo.Id(id)
	if err != nil { t.Fatal(err) }

	if result.Answer != 1 {
		t.Errorf("Differs")
	}
}



