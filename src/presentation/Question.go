package presentation

import (
	"deffish-server/src/domain"
	"deffish-server/src/helpers"
)

type Question struct {
	Id string `json:"id"`
	PDF     []byte `json:"pdf"`
	Answer  int `json:"answer"`
	Choices []string `json:"choices"`
	Tags    []string `json:"tags"`
}

func fromRequestToQuestion(question Question) domain.Question {
	var choices []domain.Choice
	for _, element := range question.Choices {
		choices = append(choices, domain.Choice{
			Content: element,
		})
	}

	var tags []domain.Tag
	for _, element := range question.Tags {
		tags = append(tags, domain.Tag{
			Name: element,
		})
	}

	return domain.Question{
		PDF: domain.PDF{
			Content: question.PDF,
		},
		Answer: question.Answer,
		Choices: choices,
		Tags: tags,
	}
}

func fromQuestionToJson(question domain.Question) Question {
	return Question{
		Id: question.Id.Value,
		PDF: question.PDF.Content,
		Answer: question.Answer,
		Choices: helpers.ChoicesToStringArray(question.Choices),
		Tags: helpers.TagsToStringArray(question.Tags),
	}
}


