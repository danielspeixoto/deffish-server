package helpers

import "deffish-server/src/aggregates"

func TextArrToStringArray(objArr []aggregates.Text) []string {
	var arrStr = make([]string, 0)
	for _, element := range objArr {
		arrStr = append(arrStr, element.Value)
	}
	return arrStr
}

