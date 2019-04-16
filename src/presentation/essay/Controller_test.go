package essay

import (
	"bytes"
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/essay"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)

func TestRouter(t *testing.T) {
	port := "5001"
	relativePath := "/essays"
	url := "http://localhost:" + port + relativePath
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	upload := essay.NewMockIUploadUseCase(ctrl)
	random := essay.NewMockIRandomUseCase(ctrl)
	byId := essay.NewMockIByIdUseCase(ctrl)
	filterByTopic := essay.NewMockIFilterByTopicUseCase(ctrl)
	comment := essay.NewMockICommentUseCase(ctrl)


	essays := Router{
		Controller: func(presenter Presenter) Controller {
			return Controller {
				UploadUseCase: upload,
				RandomUseCase:random,
				GetById:byId,
				FilterByTopicUseCase:filterByTopic,
				CommentUseCase:comment,
			}
		},
	}
	router := gin.Default()
	essays.Route(router.Group(relativePath))
	go func() {
		err := router.Run(":" + port)
		if err != nil {
			panic(err)
		}
	}()


	t.Run("Upload", func(t *testing.T) {

		upload.EXPECT().Upload(gomock.Eq(example))
		_, err := http.Post(
			url,
			"application/json",
			bytes.NewBuffer(exampleJson))

		if err != nil { panic(err) }

	})

	t.Run("random", func(t *testing.T) {

		random.EXPECT().Random(2)
		_, err := http.Get(url + "?amount=2&mode=random")
		if err != nil { panic(err) }

	})
	t.Run("Id", func(t *testing.T) {

		byId.EXPECT().
			Id(gomock.Eq(aggregates.Id{Value: "2"}))

		_, err := http.Get(url + "/2")
		if err != nil { panic(err) }

	})
	t.Run("Comment", func(t *testing.T) {
		comment.EXPECT().
			Comment(aggregates.Id{"2"}, aggregates.Comment{aggregates.Text{"A"}})

		comment , _ := json.Marshal("A")

		_, err := http.Post(
			url + "/2/comment",
			"application/json",
			bytes.NewBuffer(comment))
		if err != nil { panic(err) }
	})

	t.Run("Filter by topic", func(t *testing.T) {
		filterByTopic.EXPECT().FilterByTopic(aggregates.Id{"2"})
		_, err := http.Get(url + "?topicId=2")
		if err != nil { panic(err) }

	})
}