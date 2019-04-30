package tag

import (
	"deffish-server/src/aggregates"
	boundary "deffish-server/src/boundary/tag"
	"deffish-server/src/presentation/data"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)


type Controller struct {
	SuggestionsUseCase boundary.ISuggestionsBySubStr
	ByNameUseCase      boundary.IByNameUseCase
	UploadUseCase      boundary.IUploadUseCase
}

func (ctrl Controller) Post(c *gin.Context) {
	request := c.Request
	bodyBytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}

	var t Tag
	err = json.Unmarshal(bodyBytes, &t)
	if err != nil {
		log.Printf("request body of failed json parsing: %s", request.Body)
		panic(err)
	}

	res, err := ctrl.UploadUseCase.Upload(fromRequestToTag(t))
	if err != nil {
		panic(err)
	}
	err = json.NewEncoder(c.Writer).Encode(data.Response{
		Status: "ok",
		Data:   data.Id{res.Value},
	})
	if err != nil {
		panic(err)
	}
}

func (ctrl Controller) Get(c *gin.Context) {
	query := c.Request.URL.Query()
	hasQuestions := query.Get("hasQuestions")
	if query.Get("mode") == "suggestion" {
		q := query.Get("query")
		var res []aggregates.Tag
		var err error
		if hasQuestions != "" {
			res, err = ctrl.SuggestionsUseCase.GetSuggestionsWithQuestions(q)
		} else {
			res, err = ctrl.SuggestionsUseCase.GetSuggestions(q)
		}
		if err != nil {
			panic(err)
		}
		tags := make([]Tag, 0)
		for _, element := range res {
			tags = append(tags, fromTagToJson(element))
		}
		err = json.NewEncoder(c.Writer).Encode(data.Response{
			Status: "ok",
			Data:   tags,
		})
		if err != nil {
			panic(err)
		}

	} else {
		name := query.Get("name")
		res, err := ctrl.ByNameUseCase.ByName(name)
		if err != nil {
			panic(err)
		}
		err = json.NewEncoder(c.Writer).Encode(data.Response{
			Status: "ok",
			Data:   fromTagToJson(res),
		})
		if err != nil {
			panic(err)
		}
	}
}
