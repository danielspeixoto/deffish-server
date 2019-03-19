package essay

import (
	"bytes"
	"deffish-server/src/boundary/essay"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)

func TestController_Upload(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	uploadEssay := essay.NewMockIUploadUseCase(mockCtrl)

	controller := Controller{UploadUseCase: uploadEssay}

	uploadEssay.EXPECT().
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

	useCase := essay.NewMockIRandomUseCase(mockCtrl)

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

	useCase := essay.NewMockIRandomUseCase(mockCtrl)

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