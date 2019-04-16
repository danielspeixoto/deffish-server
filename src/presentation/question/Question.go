package question

import (
	"deffish-server/src/aggregates"
)

type Question struct {
	Id string `json:"id"`
	Image     []byte `json:"image"`
	Source string `json:"source"`
	Variant string `json:"variant"`
	Edition int `json:"edition"`
	Number int `json:"number"`
	Domain string `json:"domain"`
	Answer  int `json:"answer"`
	Tags    []string `json:"tags"`
}

func fromRequestToQuestion(question Question) aggregates.Question {
	return aggregates.Question{
		Image: aggregates.Image{
			question.Image,
		},
		Source: question.Source,
		Variant: question.Variant,
		Edition: question.Edition,
		Number: question.Number,
		Domain: question.Domain,
		Answer: question.Answer,
		Tags: question.Tags,
	}
}

func fromQuestionToJson(question aggregates.Question) Question {
	return Question{
		Id: question.Id.Value,
		Image: question.Image.Contents,
		Source: question.Source,
		Variant: question.Variant,
		Edition: question.Edition,
		Number: question.Number,
		Domain: question.Domain,
		Answer: question.Answer,
		Tags: question.Tags,
	}
}


