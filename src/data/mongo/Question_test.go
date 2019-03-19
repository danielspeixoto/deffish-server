package mongo

import (
	"bytes"
	"deffish-server/src/aggregates"
	"io/ioutil"
	"reflect"
	"strconv"
	"testing"
)


func TestQuestionOneItem(t *testing.T) {
	pdfBytes, err := ioutil.ReadFile("../res/question.pdf")
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
	t.Run("Insert One Item", func(t *testing.T) {
		id, err := questionRepo.Insert(question)
		if err != nil { t.Fatal(err) }

		t.Run("Retrieve using Id", func(t *testing.T) {
			result, err := questionRepo.Id(id)
			if err != nil { t.Fatal(err) }

			if result.Answer != 0 {
				t.Errorf("Differs")
			}
		})

		t.Run("Find All", func(t *testing.T) {
			questions, err := questionRepo.Find()
			if err != nil { t.Fatal(err) }

			question.Id = questions[0].Id
			if !reflect.DeepEqual(questions[0], question) {
				t.Fatal("Objects are different")
			}

			mongoPdfBytes := questions[0].PDF.Content
			if bytes.Compare(mongoPdfBytes, pdfBytes) != 0 {
				t.Fatal("PDF is different")
			}
		})
	})
}

func TestQuestionManyItems(t *testing.T) {
	t.Run("Insert Items", func(t *testing.T) {
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
			_, err := questionRepo.Insert(question)
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
			_, err := questionRepo.Insert(question)
			if err != nil { t.Fatal(err) }
		}

		for i := 0; i < 5; i++ {
			question := aggregates.Question{
				PDF: aggregates.PDF{
					Content: []byte{1},
				},
				Answer: i + 10,
				Tags: [] aggregates.Tag{},
			}
			_, err := questionRepo.Insert(question)
			if err != nil { t.Fatal(err) }
		}

		t.Run("Random with tags should only retrieve questions with those tags", func(t *testing.T) {
			questions, err := questionRepo.Random(100, []aggregates.Tag{{"other"}})
			if err != nil { t.Fatal(err) }

			if len(questions) != 5 {
				t.Errorf("Random should return questions with tag. " +
					"Expected: %v, Got: %v", 5, len(questions))
			}


		})

		t.Run("Random should give different results each time is run", func(t *testing.T) {
			var randomQuestions []aggregates.Question
			for i := 0; i < 10; i++ {
				questions, err := questionRepo.Random(1, []aggregates.Tag{{"other"}})
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
		})

		t.Run("Random without tags should retrieve all", func(t *testing.T) {

			questions, err := questionRepo.Random(100, []aggregates.Tag{})
			if err != nil { t.Fatal(err) }

			if len(questions) != 15 {
				t.Errorf("Random should return all questions. " +
					"Expected: %v, Got: %v", 15, len(questions))
			}
		})
	})
}





