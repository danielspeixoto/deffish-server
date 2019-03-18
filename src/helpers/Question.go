package helpers

import (
	"deffish-server/src/aggregates"
)

func TagsToStringArray(tags []aggregates.Tag) []string {
	var tagsStr = make([]string, 0)
	for _, element := range tags {
		tagsStr = append(tagsStr, element.Name)
	}
	return tagsStr
}

func ChoicesToStringArray(choices []aggregates.Choice) []string {
	var choicesStr = make([]string, 0)
	for _, element := range choices {
		choicesStr = append(choicesStr, element.Content)
	}
	return choicesStr
}
