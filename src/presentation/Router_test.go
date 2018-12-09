package presentation

import (
	"deffish-server/src/domain/gateway/mock"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"io/ioutil"
	"net/http"
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



}