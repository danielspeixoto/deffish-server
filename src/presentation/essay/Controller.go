package essay

import (
	"deffish-server/src/aggregates"
	boundary "deffish-server/src/boundary/essay"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"strconv"
)

type Controller struct {
	UploadUseCase boundary.IUploadUseCase
	RandomUseCase boundary.IRandomUseCase
	GetById boundary.IByIdUseCase
	FilterByTopicUseCase boundary.IFilterByTopicUseCase
	CommentUseCase boundary.ICommentUseCase
}

func (ctrl Controller) FilterByTopic(c *gin.Context)  {
	request := c.Request
	topicId := request.URL.Query().Get("topicId")
	if topicId != "" {
		ctrl.FilterByTopicUseCase.FilterByTopic(aggregates.Id{topicId})
	}
}

func (ctrl Controller) Comment(c *gin.Context)  {
	request := c.Request
	bodyBytes, err := ioutil.ReadAll(request.Body)
	if err != nil { panic(err) }

	var comment Comment
	err = json.Unmarshal(bodyBytes, &comment)
	if err != nil {
		log.Printf("request body of failed json parsing: %s", request.Body)
		panic(err)
	}

	id := c.Param("id")
	if id == "" { panic(errors.New("essay id is empty")) }
	ctrl.CommentUseCase.Comment(aggregates.Id{id}, fromRequestToComment(comment))
}



func (ctrl Controller) Upload(c *gin.Context)  {
	request := c.Request
	bodyBytes, err := ioutil.ReadAll(request.Body)
	if err != nil { panic(err) }

	var essay Essay
	err = json.Unmarshal(bodyBytes, &essay)
	if err != nil {
		log.Printf("request body of failed json parsing: %s", request.Body)
		panic(err)
	}

	ctrl.UploadUseCase.Upload(
		fromRequestToEssay(essay))
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

func (ctrl Controller) EssayById(c *gin.Context)  {
	id := c.Param("id")
	ctrl.GetById.Id(aggregates.Id{Value: id})
}
