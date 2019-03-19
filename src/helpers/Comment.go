package helpers

import "deffish-server/src/aggregates"

func CommentArrToStringArray(objArr []aggregates.Comment) []string {
	var arrStr = make([]string, 0)
	for _, element := range objArr {
		arrStr = append(arrStr, element.Text.Value)
	}
	return arrStr
}

