package presentation

import (
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