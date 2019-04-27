package question

import (
	"deffish-server/src/aggregates"
)

type Question struct {
	Id          string   `json:"id"`
	View        []byte   `json:"view"`
	Source      string   `json:"source"`
	Variant     string   `json:"variant"`
	Edition     int      `json:"edition"`
	Number      int      `json:"number"`
	Domain      string   `json:"domain"`
	Answer      int      `json:"answer"`
	Tags        []string `json:"tags"`
	ItemCode    string   `json:"itemCode"`
	ReferenceId string   `json:"referenceId"`
	Stage int `json:"stage"`
}

func fromQuestionsToJsonArray(questions []aggregates.Question) []Question {
	jsons := make([]Question, 0)
	for _, q := range questions {
		jsons = append(jsons, fromQuestionToJson(q))
	}
	return jsons
}

func fromRequestToQuestion(question Question) aggregates.Question {
	return aggregates.Question{
		View: aggregates.View{
			question.View,
		},
		Source:      question.Source,
		Variant:     question.Variant,
		Edition:     question.Edition,
		Number:      question.Number,
		Domain:      question.Domain,
		Answer:      question.Answer,
		Tags:        question.Tags,
		ItemCode:    question.ItemCode,
		ReferenceId: question.ReferenceId,
		Stage: question.Stage,
	}
}

func fromQuestionToJson(question aggregates.Question) Question {
	return Question{
		Id:          question.Id.Value,
		View:        question.View.Contents,
		Source:      question.Source,
		Variant:     question.Variant,
		Edition:     question.Edition,
		Number:      question.Number,
		Domain:      question.Domain,
		Answer:      question.Answer,
		Tags:        question.Tags,
		ItemCode:    question.ItemCode,
		ReferenceId: question.ReferenceId,
		Stage: question.Stage,
	}
}


