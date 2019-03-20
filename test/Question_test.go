package test

import (
	"bytes"
	"deffish-server/src/aggregates"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

var exampleJson, _ = json.Marshal(map[string]interface{}{
	"pdf":  []byte {1, 0},
	"answer": 0,
	"choices": []string{
		"A", "B", "C",
	},
	"tags": []string{
		"matematica", "enem2017",
	},
})

func TestQuestion_ManyInserted(t *testing.T) {
	port := "5000"
	relativePath := "/question"
	url := "http://localhost:" + port + relativePath

	t.Run("Upload Many", func(t *testing.T) {
		examples := make([][]byte, 0)
		for i := 0; i < 5; i++ {
			example, _ := json.Marshal(map[string]interface{}{
				"pdf":    []byte{1, 0},
				"answer": i,
				"choices": []string{
					"A", "B", "C",
				},
				"tags": []string{
					strconv.Itoa(i), "enem2017",
				},
			})
			examples = append(examples, example)
		}
		example, _ := json.Marshal(map[string]interface{}{
			"pdf":    []byte{1, 0},
			"answer": 1,
			"choices": []string{
				"A", "B", "C",
			},
			"tags": []string{
				"0",
			},
		})
		examples = append(examples, example)

		for _, example := range examples {
			resp, err := http.Post(
				url,
				"application/json",
				bytes.NewBuffer(example))

			if err != nil { panic(err) }

			expectedStatus := strconv.Itoa(http.StatusCreated)
			receivedStatus := strings.Split(resp.Status, " ")[0]
			if receivedStatus != expectedStatus {
				t.Errorf("handler returned wrong routeStatus code: got %v want %v",
					receivedStatus, expectedStatus)
			}

			jsonRespBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			var jsonResp aggregates.Response
			err = json.Unmarshal(jsonRespBytes, &jsonResp)
			if err != nil {
				panic(err)
			}

			expectedResponseStatus := "ok"
			if jsonResp.Status != expectedResponseStatus {
				t.Errorf("wrong routeStatus code: got %v want %v",
					jsonResp.Status, expectedResponseStatus)
			}
		}

		t.Run("Random items with more than one tag", func(t *testing.T) {
			resp, err := http.Get(url + "?mode=random&amount=2&tags[]=enem2017&tags[]=0")
			if err != nil { panic(err) }

			body, err := ioutil.ReadAll(resp.Body)

			var response aggregates.Response
			err = json.Unmarshal(body, &response)
			if err != nil { panic(err) }

			assert.Equal(t, "ok", response.Status)

			arr := response.Data.([]interface{})
			assert.Equal(t, len(arr), 1)
			assert.Equal(t, float64(0), arr[0].(map[string]interface{})["answer"])
		})
	})

}
