package topic

import (
	"bytes"
	"deffish-server/src/boundary/topic"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)

func TestController_Upload(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	uploadTopic := topic.NewMockIUploadUseCase(mockCtrl)

	controller := Controller{UploadUseCase: uploadTopic}

	uploadTopic.EXPECT().
		Upload(gomock.Eq(example))

	request, err := http.NewRequest(
		"POST", "/upload", bytes.NewBuffer(exampleJson))
	if err != nil { panic(err) }

	controller.Upload(&gin.Context{
		Request:request,
	})
}

func TestController_Random(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	useCase := topic.NewMockIRandomUseCase(mockCtrl)

	controller := Controller{RandomUseCase: useCase}

	useCase.EXPECT().
		Random(2)

	request, err := http.NewRequest(
		"GET", "?amount=2", nil)
	if err != nil { panic(err) }

	controller.Random(&gin.Context{
		Request:request,
	})
}

func TestController_RandomEmptyParams(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	useCase := topic.NewMockIRandomUseCase(mockCtrl)

	controller := Controller{RandomUseCase: useCase}

	useCase.EXPECT().
		Random(2)

	request, err := http.NewRequest(
		"GET", "/all", nil)
	if err != nil { panic(err) }

	controller.Random(&gin.Context{
		Request:request,
	})
}