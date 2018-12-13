package presentation

import (
	"github.com/pkg/errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestOnQuestionUploaded(t *testing.T) {
	recorder := httptest.NewRecorder()
	presenter := Presenter{Writer: recorder}
	presenter.OnQuestionUploaded()

	if status := recorder.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status , http.StatusCreated)
	}

	expected := `{"status":"ok"}`
	body := strings.TrimRight(recorder.Body.String(), "\n")
	if body != expected {
		t.Errorf("handler returned unexpected body:\n got  %v want %v",
			body, expected)
	}
}

func TestStatus(t *testing.T) {
	recorder := httptest.NewRecorder()
	presenter := Presenter{Writer: recorder}
	presenter.Status("status")

	expectedStatus := http.StatusOK

	if status := recorder.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status , expectedStatus)
	}

	expected := `{"status":"status"}`
	body := strings.TrimRight(recorder.Body.String(), "\n")
	if body != expected {
		t.Errorf("handler returned unexpected body:\n got  %v want %v",
			body, expected)
	}
}

func TestError(t *testing.T) {
	recorder := httptest.NewRecorder()
	presenter := Presenter{Writer: recorder}
	presenter.OnError(errors.New("an error explained"))

	expectedStatus := http.StatusInternalServerError

	if status := recorder.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status , expectedStatus)
	}

	expected := `{"status":"error"}`
	body := strings.TrimRight(recorder.Body.String(), "\n")
	if body != expected {
		t.Errorf("handler returned unexpected body:\n got  %v want %v",
			body, expected)
	}
}