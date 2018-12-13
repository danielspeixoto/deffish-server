package presentation

import (
	"deffish-server/src/domain"
	"deffish-server/src/domain/boundary"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Controller struct {
	UploadQuestionUseCase boundary.IUploadQuestionUseCase
	RandomQuestionUseCase boundary.IRandomQuestionUseCase
	StatusUseCase         boundary.IStatusUseCase
}

func (ctrl Controller) Upload(request *http.Request) {
	bodyBytes, err := ioutil.ReadAll(request.Body)
	if err != nil { panic(err) }

	var question Question
	err = json.Unmarshal(bodyBytes, &question)
	if err != nil { panic(err) }

	ctrl.UploadQuestionUseCase.Upload(
		fromRequestToQuestion(question))
}

func (ctrl Controller) Status(request *http.Request) {
	ctrl.StatusUseCase.Status()
}

func (ctrl Controller) Random(request *http.Request) {
	params := request.URL.Query()

	amountParam, ok := params["amount"]
	if !ok || len(amountParam[0]) < 1 {
		panic("Url param 'amount' is missing")
	}

	tagsParam, ok := params["tags[]"]
	if !ok || len(tagsParam[0]) < 1 {
		panic("Url param 'tags' is missing")
	}

	var tags []domain.Tag
	for _, element := range tagsParam {
		tags = append(tags, domain.Tag{
			Name: element,
		})
	}

	amount, err := strconv.Atoi(amountParam[0])
	if err != nil {
		panic(err)
	}

	ctrl.RandomQuestionUseCase.Random(amount, tags)
}

