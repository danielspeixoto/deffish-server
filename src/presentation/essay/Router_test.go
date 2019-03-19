package essay

import (
	"bytes"
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/essay"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
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
	_ := router.Run(":" + port)


	t.Run("Upload", func(t *testing.T) {

		upload.EXPECT().Upload(gomock.Eq(example))

		_, err := http.Post(
			url,
			"application/json",
			bytes.NewBuffer(exampleJson))

		if err != nil { panic(err) }

	})

	t.Run("Random", func(t *testing.T) {
		random.EXPECT().Random(2)

		_, err := http.Get("http://localhost:" + port + "/essays?amount=2")
		if err != nil { panic(err) }
		body, err := ioutil.ReadAll(resp.Body)

		var response aggregates.Response
		err = json.Unmarshal(body, &response)
		if err != nil { panic(err) }
		if response.Status != "ok" {
			t.Fail()
		}

		arr := response.Data.([]interface{})
		if arr[0].(map[string]interface{})["id"] != "1" {
			t.Fail()
		}
	})
	t.Run("Id", func(t *testing.T) {
		id := aggregates.Id{Value: "2"}

		repo.EXPECT().
			Id(gomock.Eq(aggregates.Id{Value: "2"})).
			Return(aggregates.Essay{Id: id}, nil)

		resp, err := http.Get("http://localhost:" + port + "/essays/2")
		if err != nil { panic(err) }
		body, err := ioutil.ReadAll(resp.Body)

		var response aggregates.Response
		err = json.Unmarshal(body, &response)
		if err != nil { panic(err) }
		if response.Status != "ok" {
			t.Fail()
		}

		if response.Data.(interface{}).(map[string]interface{})["id"] != "2" {
			t.Fail()
		}
	})
	t.Run("Comment", func(t *testing.T) {
		repo.EXPECT().
			Comment(aggregates.Id{"2"}, aggregates.Comment{aggregates.Text{"A"}})

		comment , _ := json.Marshal("A")

		resp, err := http.Post(
			"http://localhost:" + port + "/essays/2/comment",
			"application/json",
			bytes.NewBuffer(comment))

		expectedStatus := strconv.Itoa(http.StatusCreated)
		receivedStatus := strings.Split(resp.Status, " ")[0]
		if receivedStatus != expectedStatus {
			t.Errorf("handler returned wrong routeStatus code: got %v want %v",
				receivedStatus , expectedStatus)
		}

		jsonRespBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil { panic(err) }

		var jsonResp aggregates.Response
		err = json.Unmarshal(jsonRespBytes, &jsonResp)
		if err != nil { panic(err) }

		expectedResponseStatus := "ok"
		if jsonResp.Status != expectedResponseStatus {
			t.Errorf("wrong routeStatus code: got %v want %v",
				jsonResp.Status , expectedResponseStatus)
		}
	})
}