package helpers

import "deffish-server/src/domain"

func TagsToStringArray(tags []domain.Tag) []string {
	var tagsStr = []string{}
	for _, element := range tags {
		tagsStr = append(tagsStr, element.Name)
	}
	return tagsStr
}

func ChoicesToStringArray(choices []domain.Choice) []string {
	var choicesStr []string
	for _, element := range choices {
		choicesStr = append(choicesStr, element.Content)
	}
	return choicesStr
}
