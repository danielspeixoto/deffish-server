package helpers

import "deffish-server/src/aggregates"

func CommentArrToStringArray(objArr []aggregates.Comment) []string {
	var arrStr = make([]string, 0)
	for _, element := range objArr {
		arrStr = append(arrStr, element.Text.Value)
	}
	return arrStr
}


func StringArrToCommentArr(arr [] string) []aggregates.Comment {
	var arrStr = make([]aggregates.Comment, 0)
	for _, element := range arr {
		arrStr = append(arrStr, aggregates.Comment{
			aggregates.Text{element},
		})
	}
	return arrStr
}

