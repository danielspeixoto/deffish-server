package question

import (
	"deffish-server/src/aggregates"
)

type Question struct {
	Id      string   `json:"id"`
	PDF     []byte   `json:"pdf"`
	Source  string   `json:"source"`
	Variant string   `json:"variant"`
	Edition int      `json:"edition"`
	Number  int      `json:"number"`
	Domain  string   `json:"domain"`
	Answer  int      `json:"answer"`
	Tags    []string `json:"tags"`
}

func fromRequestToQuestion(question Question) aggregates.Question {
	return aggregates.Question{
		PDF: aggregates.PDF{
			question.PDF,
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
		Id:      question.Id.Value,
		PDF:     question.PDF.Contents,
		Source:  question.Source,
		Variant: question.Variant,
		Edition: question.Edition,
		Number:  question.Number,
		Domain:  question.Domain,
		Answer:  question.Answer,
		Tags:    question.Tags,
	}
}


