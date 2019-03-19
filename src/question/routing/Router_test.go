package routing

import (
	"bytes"
	"deffish-server/src/aggregates"
	"deffish-server/src/question"
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

func runServer(repo question.IRepository, port string) {
	questions := NewRouter(
		repo,
	)
	router := gin.Default()
	questions.Route(router.Group("/questions"))
	err := router.Run(":" + port)
	if err != nil {
		panic(err)
	}
}

func TestRouter(t *testing.T) {
	port := "5001"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := question.NewMockIRepository(ctrl)
	go runServer(repo, port)

	question := aggregates.Question{
		PDF: aggregates.PDF{
			Content: []byte {1, 0},
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


	t.Run("Upload", func(t *testing.T) {

		repo.EXPECT().
			Insert(gomock.Eq(question))

		body, err := json.Marshal(map[string]interface{}{
			"pdf":  []byte {1, 0},
			"answer": 0,
			"choices": []string{
				"A", "B", "C",
			},
			"tags": []string{
				"matematica", "enem2017",
			},
		})
		if err != nil { panic(err) }
		resp, err := http.Post(
			"http://localhost:" + port + "/questions",
			"application/json",
			bytes.NewBuffer(body))

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
			Insert(gomock.Eq(question)).
			Return(aggregates.Id{}, errors.New("an error"))

		body, err := json.Marshal(map[string]interface{}{
			"pdf":  []byte {1, 0},
			"answer": 0,
			"choices": []string{
				"A", "B", "C",
			},
			"tags": []string{
				"matematica", "enem2017",
			},
		})
		if err != nil { panic(err) }
		resp, err := http.Post(
			"http://localhost:" + port + "/questions/",
			"application/json",
			bytes.NewBuffer(body))

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
			Random(2, gomock.Eq([]aggregates.Tag{{Name: "enem"}, {Name: "matematica"}})).
			Return([]aggregates.Question{
				{
					Id: aggregates.Id{Value: "1"},
				},
				{
					Id: aggregates.Id{Value: "2"},
				},
			}, nil)

		resp, err := http.Get("http://localhost:" + port + "/questions?amount=2&tags[]=enem&tags[]=matematica")
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
			Return(aggregates.Question{Id: id}, nil)

		resp, err := http.Get("http://localhost:" + port + "/questions/2")
		if err != nil { panic(err) }
		body, err := ioutil.ReadAll(resp.Body)

		var response aggregates.Response
		err = json.Unmarshal(body, &response)
		if err != nil { panic(err) }
		if response.Status != "ok" {
			t.Fail()
		}

		question := response.Data.(interface{})
		if question.(map[string]interface{})["id"] != "2" {
			t.Fail()
		}
	})
}