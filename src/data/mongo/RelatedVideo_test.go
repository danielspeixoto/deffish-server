package mongo

import (
	"deffish-server/src/aggregates"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"strconv"
	"testing"
)

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
			_, err := insert(relatedVideosRepo.collection, RelatedVideo{
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
		_, err = insert(relatedVideosRepo.collection, RelatedVideo{
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
			videos, err := relatedVideosRepo.FilterByQuestion(aggregates.Id{mockId}, 0, 100)
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
			videos, err := relatedVideosRepo.FilterByQuestion(aggregates.Id{mockId}, 2, 2)
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
