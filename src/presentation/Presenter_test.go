package presentation

import (
	"deffish-server/src/domain"
	"github.com/pkg/errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPresenter_OnQuestionUploaded(t *testing.T) {
	recorder := httptest.NewRecorder()
	presenter := Presenter{Writer: recorder}
	presenter.OnQuestionUploaded()

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

func TestPresenter_Status(t *testing.T) {
	recorder := httptest.NewRecorder()
	presenter := Presenter{Writer: recorder}
	presenter.Status("status")

	expectedStatus := http.StatusOK

	if status := recorder.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status , expectedStatus)
	}

	expected := `{"status":"status","data":null}`
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

func TestPresenter_OnQuestionReceived(t *testing.T) {
	recorder := httptest.NewRecorder()
	presenter := Presenter{Writer: recorder}

	presenter.OnQuestionReceived([]domain.Question{
		{
			Id: domain.Id{Value: "1"},
			PDF: domain.PDF{Content: []byte{1,0}},
			Answer: 0,
		},
		{
			Id: domain.Id{Value: "2"},
			PDF: domain.PDF{Content: []byte{0,1}},
			Answer: 1,
		},
	})

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status , http.StatusOK)
	}

	expected := `{"status":"ok","data":[{"id":"1","pdf":"AQA=","answer":0,"choices":[],"tags":[]},{"id":"2","pdf":"AAE=","answer":1,"choices":[],"tags":[]}]}`
	body := strings.TrimRight(recorder.Body.String(), "\n")
	if body != expected {
		t.Errorf("handler returned unexpected body:\n got  %v want %v",
			body, expected)
	}
}