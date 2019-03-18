package presentation

import (
	"deffish-server/src/aggregates"
	"deffish-server/src/helpers"
)

type Question struct {
	Id string `json:"id"`
	PDF     []byte `json:"pdf"`
	Answer  int `json:"answer"`
	Choices []string `json:"choices"`
	Tags    []string `json:"tags"`
}

func fromRequestToQuestion(question Question) aggregates.Question {
	var choices []aggregates.Choice
	for _, element := range question.Choices {
		choices = append(choices, aggregates.Choice{
			Content: element,
		})
	}

	var tags []aggregates.Tag
	for _, element := range question.Tags {
		tags = append(tags, aggregates.Tag{
			Name: element,
		})
	}

	return aggregates.Question{
		PDF: aggregates.PDF{
			Content: question.PDF,
		},
		Answer: question.Answer,
		Choices: choices,
		Tags: tags,
	}
}

func fromQuestionToJson(question aggregates.Question) Question {
	return Question{
		Id: question.Id.Value,
		PDF: question.PDF.Content,
		Answer: question.Answer,
		Choices: helpers.ChoicesToStringArray(question.Choices),
		Tags: helpers.TagsToStringArray(question.Tags),
	}
}


