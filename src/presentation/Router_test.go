package presentation

import (
	"bytes"
	"deffish-server/src/domain"
	"deffish-server/src/domain/gateway/mock"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

func TestRouter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_gateway.NewMockIQuestionRepository(ctrl)
	port := 4000
	portStr := strconv.Itoa(port)
	go NewHandler(repo, port)

	t.Run("Status", func(t *testing.T) {
		resp, err := http.Get("http://localhost:" + portStr + "/status")
		if err != nil { panic(err) }
		body, err := ioutil.ReadAll(resp.Body)

		var response Response
		err = json.Unmarshal(body, &response)
		if err != nil { panic(err) }
		if response.Status != "ok" {
			t.Fail()
		}
	})

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
			"http://localhost:" + portStr + "/questions/",
			"application/json",
			bytes.NewBuffer(body))

		expectedStatus := strconv.Itoa(http.StatusCreated)
		receivedStatus := strings.Split(resp.Status, " ")[0]
		if receivedStatus != expectedStatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				receivedStatus , expectedStatus)
		}

		jsonRespBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil { panic(err) }

		var jsonResp Response
		err = json.Unmarshal(jsonRespBytes, &jsonResp)
		if err != nil { panic(err) }

		expectedResponseStatus := "ok"
		if jsonResp.Status != expectedResponseStatus {
			t.Errorf("wrong status code: got %v want %v",
				jsonResp.Status , expectedResponseStatus)
		}
	})
	//t.Run("UploadInvalid", func(t *testing.T) {
	//
	//	body, err := json.Marshal(map[string]interface{}{
	//		"pdf":  []byte {1, 0},
	//		"answer": 0,
	//		"choices": []int{
	//			2, 1, 3,
	//		},
	//		"tags": []string{
	//			"matematica", "enem2017",
	//		},
	//	})
	//	if err != nil { panic(err) }
	//	resp, err := http.Post(
	//		"http://localhost:" +  portStr + "/questions/",
	//		"application/json",
	//		bytes.NewBuffer(body))
	//
	//	expectedStatus := strconv.Itoa(http.StatusBadRequest)
	//	receivedStatus := strings.Split(resp.Status, " ")[0]
	//	if receivedStatus != expectedStatus {
	//		t.Errorf("handler returned wrong status code: got %v want %v",
	//			receivedStatus , expectedStatus)
	//	}
	//
	//	jsonRespBytes, err := ioutil.ReadAll(resp.Body)
	//	if err != nil { panic(err) }
	//
	//	var jsonResp Response
	//	err = json.Unmarshal(jsonRespBytes, &jsonResp)
	//	if err != nil { panic(err) }
	//
	//	expectedResponseStatus := "invalid"
	//	if jsonResp.Status != expectedResponseStatus  {
	//		t.Errorf("wrong status code: got %v want %v",
	//			jsonResp.Status , expectedResponseStatus)
	//	}
	//})
	t.Run("UploadError", func(t *testing.T) {

		repo.EXPECT().
			Insert(gomock.Eq(question)).
			Return(domain.Id{}, errors.New("an error"))

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
			"http://localhost:" + portStr + "/questions/",
			"application/json",
			bytes.NewBuffer(body))

		expectedStatus := strconv.Itoa(http.StatusInternalServerError)
		receivedStatus := strings.Split(resp.Status, " ")[0]
		if receivedStatus != expectedStatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				receivedStatus , expectedStatus)
		}

		jsonRespBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil { panic(err) }

		var jsonResp Response
		err = json.Unmarshal(jsonRespBytes, &jsonResp)
		if err != nil { panic(err) }

		expectedResponseStatus := "error"
		if jsonResp.Status != expectedResponseStatus {
			t.Errorf("wrong status code: got %v want %v",
				jsonResp.Status , expectedResponseStatus)
		}
	})
	t.Run("Random", func(t *testing.T) {
		repo.EXPECT().
			Random(2, gomock.Eq([]domain.Tag{{Name: "enem"}, {Name: "matematica"}})).
			Return([]domain.Question{
				{
					Id: domain.Id{Value: "1"},
				},
				{
					Id: domain.Id{Value: "2"},
				},
			}, nil)

		resp, err := http.Get("http://localhost:" + portStr + "/questions?amount=2&tags[]=enem&tags[]=matematica")
		if err != nil { panic(err) }
		body, err := ioutil.ReadAll(resp.Body)

		var response Response
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
		id := domain.Id{Value: "2"}

		repo.EXPECT().
			Id(gomock.Eq(domain.Id{Value: "2"})).
			Return(domain.Question{Id: id}, nil)

		resp, err := http.Get("http://localhost:" + portStr + "/questions/2")
		if err != nil { panic(err) }
		body, err := ioutil.ReadAll(resp.Body)

		var response Response
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