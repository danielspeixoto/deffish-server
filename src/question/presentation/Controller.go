package presentation

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/question/domain"
	"deffish-server/src/status"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Controller struct {
	UploadUseCase domain.IUploadUseCase
	RandomUseCase domain.IRandomUseCase
	StatusUseCase         status.IInfoUseCase
	GetById domain.IByIdUseCase
}

func (ctrl Controller) Upload(request *http.Request) {
	bodyBytes, err := ioutil.ReadAll(request.Body)
	if err != nil { panic(err) }

	var question Question
	err = json.Unmarshal(bodyBytes, &question)
	if err != nil {
		log.Printf("request body of failed json parsing: %s", request.Body)
		panic(err)
	}

	ctrl.UploadUseCase.Upload(
		fromRequestToQuestion(question))
}

func (ctrl Controller) Status(request *http.Request) {
	ctrl.StatusUseCase.Info()
}

func (ctrl Controller) Random(request *http.Request) {
	params := request.URL.Query()

	amountParam, ok := params["amount"]
	if !ok || len(amountParam[0]) < 1 {
		amountParam = []string{"2"}
	}

	tagsParam, ok := params["tags[]"]
	if !ok || len(tagsParam[0]) < 1 {
		tagsParam = []string{}
	}

	var tags []aggregates.Tag
	for _, element := range tagsParam {
		tags = append(tags, aggregates.Tag{
			Name: element,
		})
	}

	amount, err := strconv.Atoi(amountParam[0])
	if err != nil {
		panic(err)
	}

	ctrl.RandomUseCase.Random(amount, tags)
}

func (ctrl Controller) QuestionById(c *gin.Context)  {
	id := c.Param("id")
	ctrl.GetById.Id(aggregates.Id{Value: id})
}
