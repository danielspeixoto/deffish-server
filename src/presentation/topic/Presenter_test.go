package topic

import (
	"deffish-server/src/aggregates"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPresenter_OnTopicUploaded(t *testing.T) {
	recorder := httptest.NewRecorder()
	presenter := Presenter{Writer: recorder}
	presenter.OnUploaded()

	if status := recorder.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status , http.StatusCreated)
	}

	expected := `{"status":"ok","data":null}`
	body := strings.TrimRight(recorder.Body.String(), "\n")
	if body != expected {
		t.Errorf("handler returned unexpected body:\n got  %v want %v",
			body, expected)
	}
}

func TestPresenter_OnError(t *testing.T) {
	recorder := httptest.NewRecorder()
	presenter := Presenter{Writer: recorder}
	presenter.OnError(errors.New("an error explained"))

	expectedStatus := http.StatusInternalServerError

	if status := recorder.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status , expectedStatus)
	}

	expected := `{"status":"error","data":null}`
	body := strings.TrimRight(recorder.Body.String(), "\n")
	if body != expected {
		t.Errorf("handler returned unexpected body:\n got  %v want %v",
			body, expected)
	}
}

func TestPresenter_OnTopicsReceived(t *testing.T) {
	recorder := httptest.NewRecorder()
	presenter := Presenter{Writer: recorder}

	presenter.OnListReceived([]aggregates.Topic{
		example,
		{
			Title: aggregates.Title{
				Value: "B",
			},
			Content: []aggregates.Text{
				{"B"}, {"A"},
			},
		},
	})

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status , http.StatusOK)
	}

	result := &aggregates.Response{}
	_ = json.Unmarshal(recorder.Body.Bytes(), result)

	if result.Data.([]interface{})[0].(map[string]interface{})["title"] != example.Title.Value ||
		result.Data.([]interface{})[1].(map[string]interface{})["title"] != "B"{
		t.Fatal()
	}
}

func TestPresenter_OnTopicReceived(t *testing.T) {
	recorder := httptest.NewRecorder()
	presenter := Presenter{Writer: recorder}

	presenter.OnReceived(example)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status , http.StatusOK)
	}

	result := &aggregates.Response{}
	_ = json.Unmarshal(recorder.Body.Bytes(), result)

	if result.Data.(interface{}).(map[string]interface{})["title"] != example.Title.Value {
		t.Fatal()
	}
}