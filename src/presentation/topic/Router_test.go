package topic

import (
	"bytes"
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/topic"
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

func runServer(repo topic.IRepository, port string) {
	topics := Router{repo}
	router := gin.Default()
	topics.Route(router.Group("/topics"))
	err := router.Run(":" + port)
	if err != nil {
		panic(err)
	}
}

func TestRouter(t *testing.T) {
	port := "5001"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := topic.NewMockIRepository(ctrl)
	go runServer(repo, port)

	t.Run("Upload", func(t *testing.T) {

		repo.EXPECT().
			Insert(gomock.Eq(example))

		resp, err := http.Post(
			"http://localhost:" + port + "/topics",
			"application/json",
			bytes.NewBuffer(exampleJson))

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
	t.Run("UploadError", func(t *testing.T) {

		repo.EXPECT().
			Insert(gomock.Eq(example)).
			Return(aggregates.Id{}, errors.New("an error"))

		resp, err := http.Post(
			"http://localhost:" + port + "/topics/",
			"application/json",
			bytes.NewBuffer(exampleJson))

		expectedStatus := strconv.Itoa(http.StatusInternalServerError)
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

		expectedResponseStatus := "error"
		if jsonResp.Status != expectedResponseStatus {
			t.Errorf("wrong routeStatus code: got %v want %v",
				jsonResp.Status , expectedResponseStatus)
		}
	})
	t.Run("Random", func(t *testing.T) {
		repo.EXPECT().
			Random(2).
			Return([]aggregates.Topic{
				{
					Id: aggregates.Id{Value: "1"},
				},
				{
					Id: aggregates.Id{Value: "2"},
				},
			}, nil)

		resp, err := http.Get("http://localhost:" + port + "/topics?amount=2")
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
			Return(aggregates.Topic{Id: id}, nil)

		resp, err := http.Get("http://localhost:" + port + "/topics/2")
		if err != nil { panic(err) }
		body, err := ioutil.ReadAll(resp.Body)

		var response aggregates.Response
		err = json.Unmarshal(body, &response)
		if err != nil { panic(err) }
		if response.Status != "ok" {
			t.Fail()
		}

		topic := response.Data.(interface{})
		if topic.(map[string]interface{})["id"] != "2" {
			t.Fail()
		}
	})
}