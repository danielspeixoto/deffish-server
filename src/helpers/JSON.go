package helpers

import (
	"deffish-server/src/presentation/data"
	"encoding/json"
	"net/http"
)

func FromResponseToMap(res *http.Response) data.Response {
	var bodyBytes []byte
	_, err := res.Body.Read(bodyBytes)
	if err != nil { panic(err) }

	var jsonMap data.Response
	err = json.Unmarshal(bodyBytes, &jsonMap)
	if err != nil { panic(err) }
	return jsonMap
}
