package question

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/question"
	"deffish-server/src/presentation/data"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"strconv"
)

type Controller struct {
	UploadUseCase     question.IUploadUseCase
	RandomTagsUseCase question.IRandomByTagsUseCase
	GetByIdUseCase    question.IByIdUseCase
	VideosUseCase     question.IGetRelatedVideos
	AddTagUseCase     question.IAddTagUseCase
}

func (ctrl Controller) Upload(c *gin.Context) {
	request := c.Request
	bodyBytes, err := ioutil.ReadAll(request.Body)
	if err != nil { panic(err) }

	var q Question
	err = json.Unmarshal(bodyBytes, &q)
	if err != nil {
		log.Printf("request body of failed json parsing: %s", string(bodyBytes))
		panic(err)
	}

	id, err := ctrl.UploadUseCase.Upload(
		fromRequestToQuestion(q))
	if err != nil {
		panic(err)
	}
	err = json.NewEncoder(c.Writer).Encode(data.Response{
		Status: "ok",
		Data:   data.Id{id.Value},
	})
	if err != nil {
		panic(err)
	}
}

func (ctrl Controller) RandomByTags(c *gin.Context) {
	request := c.Request
	params := request.URL.Query()

	amountParam, ok := params["amount"]
	if !ok || len(amountParam[0]) < 1 {
		amountParam = []string{"2"}
	}

	tagsParam, ok := params["tags[]"]
	if !ok || len(tagsParam[0]) < 1 {
		tagsParam = []string{}
	}

	amount, err := strconv.Atoi(amountParam[0])
	if err != nil {
		panic(err)
	}

	print(tagsParam)

	questions, err := ctrl.RandomTagsUseCase.Random(amount, tagsParam)
	if err != nil {
		panic(err)
	}
	err = json.NewEncoder(c.Writer).Encode(data.Response{
		Status: "ok",
		Data:   fromQuestionsToJsonArray(questions),
	})
	if err != nil {
		panic(err)
	}
}

func (ctrl Controller) QuestionById(c *gin.Context)  {
	id := c.Param("id")
	q, err := ctrl.GetByIdUseCase.Id(aggregates.Id{Value: id})
	if err != nil {
		panic(err)
	}
	err = json.NewEncoder(c.Writer).Encode(data.Response{
		Status: "ok",
		Data:   fromQuestionToJson(q),
	})
	if err != nil {
		panic(err)
	}
}

func (ctrl Controller) RelatedVideos(c *gin.Context) {
	id := c.Param("id")
	query := c.Request.URL.Query()
	amount := query.Get("amount")
	start := query.Get("start")
	if amount == "" || start == "" || id == "" {
		panic(errors.New("not all params supplied"))
	}
	startInt, err := strconv.Atoi(start)
	if err != nil {
		panic(err)
	}
	amountInt, err := strconv.Atoi(amount)
	if err != nil {
		panic(err)
	}
	v, err := ctrl.VideosUseCase.GetRelatedVideos(aggregates.Id{id}, startInt, amountInt)
	if err != nil {
		panic(err)
	}
	err = json.NewEncoder(c.Writer).Encode(data.Response{
		Status: "ok",
		Data:   fromRelatedVideosToJsonArray(v),
	})
	if err != nil {
		panic(err)
	}
}

func (ctrl Controller) AddTag(c *gin.Context) {
	id := c.Param("id")

	request := c.Request
	bodyBytes, err := ioutil.ReadAll(request.Body)
	if err != nil { panic(err) }
	var t struct {
		Name string `json:"name"`
	}
	err = json.Unmarshal(bodyBytes, &t)
	if err != nil {
		log.Printf("request body of failed json parsing: %s", string(bodyBytes))
		panic(err)
	}

	err = ctrl.AddTagUseCase.Add(aggregates.Id{id}, t.Name)
	if err != nil {
		panic(err)
	}
	err = json.NewEncoder(c.Writer).Encode(data.Response{
		Status: "ok",
	})
	if err != nil {
		panic(err)
	}
}