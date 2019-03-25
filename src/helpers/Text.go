package helpers

import (
	"deffish-server/src/aggregates"
)

func TextArrToStringArray(objArr []aggregates.Text) []string {
	var arrStr = make([]string, 0)
	for _, element := range objArr {
		arrStr = append(arrStr, element.Value)
	}

	return arrStr
}

func StringArrToTextArr(arr [] string) []aggregates.Text {
	var arrStr = make([]aggregates.Text, 0)
	for _, element := range arr {
		arrStr = append(arrStr, aggregates.Text{
			element,
		})
	}
	return arrStr
}