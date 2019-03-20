package question

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/domain/question"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Controller struct {
	UploadUseCase question.IUploadUseCase
	RandomUseCase question.IRandomUseCase
	GetById question.IByIdUseCase
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