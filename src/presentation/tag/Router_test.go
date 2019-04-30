package tag

import (
	"bytes"
	"deffish-server/src/aggregates"
	"deffish-server/src/boundary/tag"
	"deffish-server/src/presentation/data"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"io/ioutil"
	"net/http"
	"testing"
)

var example = Tag{
	Name: "mytag",
}

var exampleJson, _ = json.Marshal(example)

func TestRouter(t *testing.T) {
	port := "5001"
	relativePath := "/tag"
	url := "http://localhost:" + port + relativePath
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	upload := tag.NewMockIUploadUseCase(ctrl)
	byName := tag.NewMockIByNameUseCase(ctrl)
	sugg := tag.NewMockISuggestionsBySubStr(ctrl)

	tags := Router{
		Controller{
			UploadUseCase:      upload,
			ByNameUseCase:      byName,
			SuggestionsUseCase: sugg,
		},
	}
	router := gin.Default()
	tags.Route(router.Group(relativePath))
	go func() {
		err := router.Run(":" + port)
		if err != nil {
			panic(err)
		}
	}()

	t.Run("Upload", func(t *testing.T) {
		upload.EXPECT().
			Upload(gomock.Any()).
			Return(aggregates.Id{"123"}, nil)
		res, err := http.Post(
			url,
			"application/json",
			bytes.NewBuffer(exampleJson))

		if err != nil {
			panic(err)
		}

		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		var jsonMap data.Response
		err = json.Unmarshal(bodyBytes, &jsonMap)
		if err != nil {
			panic(err)
		}

		jsonData := jsonMap.Data.(map[string]interface{})

		if jsonData["id"] != "123" {
			t.Errorf("Did not return the correct id")
		}
	})

	t.Run("Get By Name", func(t *testing.T) {
		byName.EXPECT().
			ByName(gomock.Any()).
			Return(aggregates.Tag{
					Id: aggregates.Id{
						"321",
					},
					Name: "mytag",
				}, nil)

		res, err := http.Get(
			url + "?name=mytag",
		)

		if err != nil {
			panic(err)
		}

		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		var jsonMap data.Response
		err = json.Unmarshal(bodyBytes, &jsonMap)
		if err != nil {
			panic(err)
		}

		jsonData := jsonMap.Data.(map[string]interface{})

		if jsonData["name"] != "mytag" {
			t.Errorf("Name is Different")
		}
		if jsonData["id"] != "321" {
			t.Errorf("Id is Different")
		}
	})

	t.Run("Suggestions", func(t *testing.T) {
		sugg.EXPECT().
			GetSuggestions("abc").
			Return([]aggregates.Tag{
				{
					Id: aggregates.Id{
						"1",
					},
					Name: "abc",
				},
				{
					Id: aggregates.Id{
						"2",
					},
					Name: "dcabc",
				},
				{
					Id: aggregates.Id{
						"3",
					},
					Name: "aaaabc",
				},
			}, nil)
		res, err := http.Get(
			url + "?mode=suggestion&query=abc")
		if err != nil {
			panic(err)
		}

		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		var jsonMap data.Response
		err = json.Unmarshal(bodyBytes, &jsonMap)
		if err != nil {
			panic(err)
		}

		jsonData := jsonMap.Data.([]interface{})

		if len(jsonData) != 3 {
			t.Errorf("Size is Different")
		}
		if jsonData[0].(map[string]interface{})["id"] != "1" {
			t.Errorf("Id is Different")
		}
		if jsonData[1].(map[string]interface{})["name"] != "dcabc" {
			t.Errorf("Name is Different")
		}
	})

	t.Run("Suggestions only questions with tags", func(t *testing.T) {
		sugg.EXPECT().
			GetSuggestionsWithQuestions("abc").
			Return([]aggregates.Tag{
				{
					Id: aggregates.Id{
						"1",
					},
					Name: "abc",
				},
				{
					Id: aggregates.Id{
						"2",
					},
					Name: "dcabc",
				},
				{
					Id: aggregates.Id{
						"3",
					},
					Name: "aaaabc",
				},
			}, nil)
		res, err := http.Get(
			url + "?mode=suggestion&query=abc&hasTags=1")
		if err != nil {
			panic(err)
		}

		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		var jsonMap data.Response
		err = json.Unmarshal(bodyBytes, &jsonMap)
		if err != nil {
			panic(err)
		}

		jsonData := jsonMap.Data.([]interface{})

		if len(jsonData) != 3 {
			t.Errorf("Size is Different")
		}
		if jsonData[0].(map[string]interface{})["id"] != "1" {
			t.Errorf("Id is Different")
		}
		if jsonData[1].(map[string]interface{})["name"] != "dcabc" {
			t.Errorf("Name is Different")
		}
	})
}
