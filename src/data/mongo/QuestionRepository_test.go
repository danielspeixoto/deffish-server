package mongo

import (
	"deffish-server/src/aggregates"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"reflect"
	"strconv"
	"testing"
)

var testQuestion = aggregates.Question{
	View: aggregates.View{
		[]byte{0, 0, 1},
	},
	Source:  "enem",
	Variant: "AMARELO",
	Edition: 2017,
	Number:  3,
	Domain:  "Linguagens",
	Answer:  1,
	Tags:    []string{"a", "b"},
}

func TestQuestionOneItem(t *testing.T) {
	t.Run("Insert One Item", func(t *testing.T) {
		id, err := questionRepo.Insert(testQuestion)
		if err != nil { t.Fatal(err) }

		t.Run("Retrieve using Id", func(t *testing.T) {
			result, err := questionRepo.Id(id)
			if err != nil { t.Fatal(err) }

			if reflect.DeepEqual(result, testQuestion) {
				t.Errorf("Differs")
			}
		})

		t.Run("FindAll", func(t *testing.T) {
			questions, err := questionRepo.Find()
			if err != nil { t.Fatal(err) }

			testQuestion.Id = questions[0].Id
			if !reflect.DeepEqual(questions[0], testQuestion) {
				t.Fatal("Objects are different")
			}
		})

		t.Run("Add Tag", func(t *testing.T) {
			err := questionRepo.Add(id, "aTag")
			if err != nil { t.Fatal(err) }

			result, err := questionRepo.Id(id)
			if err != nil { t.Fatal(err) }

			if !reflect.DeepEqual(result.Tags, []string{"a", "b", "aTag"}) {
				t.Fatal()
			}
		})
	})
}

func TestQuestionManyItems(t *testing.T) {
	t.Run("Insert Items", func(t *testing.T) {
		for i := 0; i < 5; i++ {
			question := aggregates.Question{
				Source:  "enem",
				Variant: "AMARELO",
				Edition: 2017,
				Number:  3,
				Domain:  strconv.Itoa(i),
				Answer:  i,
				Tags:    []string {
					strconv.Itoa(i),
				},
			}
			_, err := questionRepo.Insert(question)
			if err != nil { t.Fatal(err) }
		}

		for i := 0; i < 5; i++ {
			question := aggregates.Question{
				View:    aggregates.View{},
				Source:  "enem",
				Variant: "AMARELO",
				Edition: 2017,
				Number:  3,
				Domain:  "other",
				Tags:    []string {
					strconv.Itoa(i),
					"other",
				},
				Answer: i + 5,
			}
			_, err := questionRepo.Insert(question)
			if err != nil { t.Fatal(err) }
		}

		for i := 0; i < 5; i++ {
			question := aggregates.Question{
				View:    aggregates.View{},
				Source:  "enem",
				Variant: "AMARELO",
				Edition: 2017,
				Number:  3,
				Domain:  "none",
				Tags:    []string {},
				Answer:  i + 10,
			}
			_, err := questionRepo.Insert(question)
			if err != nil { t.Fatal(err) }
		}

		t.Run("random with domain should only retrieve questionsCollection with this domain", func(t *testing.T) {
			questions, err := questionRepo.RandomByDomain(100, "other")
			if err != nil { t.Fatal(err) }

			if len(questions) != 5 {
				t.Errorf("random should return questionsCollection with domain. " +
					"Expected: %v, Got: %v", 5, len(questions))
			}


		})

		t.Run("random should give different results each time is run", func(t *testing.T) {
			var randomQuestions []aggregates.Question
			for i := 0; i < 10; i++ {
				questions, err := questionRepo.RandomByDomain(1, "other")
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

		t.Run("RandomByTags with tags should only retrieve questionsCollection with those tags", func(t *testing.T) {
			questions, err := questionRepo.RandomByTags(100, []string{"other"})
			if err != nil { t.Fatal(err) }

			if len(questions) != 5 {
				t.Errorf("RandomByDomain should return questionsCollection with tag. " +
					"Expected: %v, Got: %v", 5, len(questions))
			}


		})

		t.Run("RandomByTags without tags should retrieve all", func(t *testing.T) {

			questions, err := questionRepo.RandomByTags(100, []string{})
			if err != nil { t.Fatal(err) }

			if len(questions) != 15 {
				t.Errorf("RandomByDomain should return all questionsCollection. " +
					"Expected: %v, Got: %v", 15, len(questions))
			}
		})
	})
}

func TestRelatedVideosManyItems(t *testing.T) {
	mockId := "5cbe17832910520c3d84bcd1"
	t.Run("Insert Many", func(t *testing.T) {
		objId, err := primitive.ObjectIDFromHex(mockId)
		if err != nil {
			panic(err)
		}
		total := 10
		for i := 0; i < total; i++ {
			retrieve := i
			if i <= 5 {
				retrieve = 5 - i
			}
			_, err := insert(questionRepo.relatedVideosCollection, RelatedVideo{
				QuestionId:        objId,
				RetrievalPosition: retrieve,
				Title:             strconv.Itoa(retrieve),
				Channel: Channel{
					Title: "c",
					Id:    "1",
				},
				Thumbnails: Thumbnails{
					High:    "a",
					Default: "b",
					Medium:  "c",
				},
				Description: "desc",
				VideoId:     "1",
			})
			if err != nil {
				panic(err)
			}
		}
		_, err = insert(questionRepo.relatedVideosCollection, RelatedVideo{
			QuestionId: primitive.ObjectID{
				1, 1, 1, 1,
				1, 1, 1, 1,
				1, 1, 1, 1,
			},
			RetrievalPosition: 1,
			Title:             "other",
			Channel: Channel{
				Title: "c",
				Id:    "1",
			},
			Thumbnails: Thumbnails{
				High:    "a",
				Default: "b",
				Medium:  "c",
			},
			Description: "desc",
			VideoId:     "1",
		})
		if err != nil {
			panic(err)
		}

		t.Run("Filter By Question", func(t *testing.T) {
			videos, err := questionRepo.GetRelatedVideos(aggregates.Id{mockId}, 0, 100)
			if err != nil {
				panic(err)
			}
			if len(videos) != total {
				t.Fatalf("Not filtering")
			}
			for i := 0; i < total; i++ {
				if videos[i].Title != strconv.Itoa(i) {
					t.Fatalf("Not ordered by retrieve list")
				}
			}
		})

		t.Run("Filter By Question using Pagination", func(t *testing.T) {
			videos, err := questionRepo.GetRelatedVideos(aggregates.Id{mockId}, 2, 2)
			if err != nil {
				panic(err)
			}
			if len(videos) != 2 {
				t.Fail()
			}
			for i := 0; i < 2; i++ {
				if videos[i].Title != strconv.Itoa(i + 2) {
					t.Fail()
				}
			}
		})
	})
}






