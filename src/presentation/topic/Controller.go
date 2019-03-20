package topic

import (
	"deffish-server/src/aggregates"
	boundary "deffish-server/src/boundary/topic"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"strconv"
)

type Controller struct {
	UploadUseCase boundary.IUploadUseCase
	RandomUseCase boundary.IRandomUseCase
	GetById boundary.IByIdUseCase
}

func (ctrl Controller) Upload(c *gin.Context)  {
	request := c.Request
	bodyBytes, err := ioutil.ReadAll(request.Body)
	if err != nil { panic(err) }

	var topic Topic
	err = json.Unmarshal(bodyBytes, &topic)
	if err != nil {
		log.Printf("request body of failed json parsing: %s", request.Body)
		panic(err)
	}

	ctrl.UploadUseCase.Upload(
		fromRequestToTopic(topic))
}

func (ctrl Controller) Random(c *gin.Context)  {
	request := c.Request
	params := request.URL.Query()

	amountParam, ok := params["amount"]
	if !ok || len(amountParam[0]) < 1 {
		amountParam = []string{"2"}
	}

	amount, err := strconv.Atoi(amountParam[0])
	if err != nil {
		panic(err)
	}

	ctrl.RandomUseCase.Random(amount)
}

func (ctrl Controller) TopicById(c *gin.Context)  {
	id := c.Param("id")
	ctrl.GetById.Id(aggregates.Id{Value: id})
}
