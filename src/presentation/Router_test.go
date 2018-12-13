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

func TestStatusRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_gateway.NewMockIQuestionRepository(ctrl)
	go NewHandler(repo, 5000)

	resp, err := http.Get("http://localhost:5000/status")
	if err != nil { panic(err) }
	body, err := ioutil.ReadAll(resp.Body)

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil { panic(err) }
	if response.Status != "ok" {
		t.Fail()
	}
}

func TestUploadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_gateway.NewMockIQuestionRepository(ctrl)
	go NewHandler(repo, 5000)

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
		"http://localhost:5000/upload",
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
}

func TestUploadRequestInvalid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_gateway.NewMockIQuestionRepository(ctrl)
	go NewHandler(repo, 5000)

	body, err := json.Marshal(map[string]interface{}{
		"pdf":  []byte {1, 0},
		"answer": 0,
		"choices": []int{
			2, 1, 3,
		},
		"tags": []string{
			"matematica", "enem2017",
		},
	})
	if err != nil { panic(err) }
	resp, err := http.Post(
		"http://localhost:5000/upload",
		"application/json",
		bytes.NewBuffer(body))

	expectedStatus := strconv.Itoa(http.StatusBadRequest)
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

	expectedResponseStatus := "invalid"
	if jsonResp.Status != expectedResponseStatus  {
		t.Errorf("wrong status code: got %v want %v",
			jsonResp.Status , expectedResponseStatus)
	}
}

func TestUploadRequestError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_gateway.NewMockIQuestionRepository(ctrl)
	go NewHandler(repo, 5000)

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
		"http://localhost:5000/upload",
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
}