package presentation

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

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