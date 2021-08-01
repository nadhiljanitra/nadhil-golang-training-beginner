package healthcheck

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(health)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	responseJson := healthCheckResponse{}
	err = json.Unmarshal(rr.Body.Bytes(), &responseJson)
	if err != nil {
		t.Errorf("Failed on unmarshalling response")
	}

	expected := healthCheckResponse{
		Status: "healthy",
	}
	if responseJson != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", responseJson, expected)
	}

	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("content type header does not match: got %v want %v",
			ctype, "application/json")
	}
}

func TestHelloWorldHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello-world", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloWorld)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	responseJson := helloWorldResponse{}
	err = json.Unmarshal(rr.Body.Bytes(), &responseJson)
	if err != nil {
		t.Errorf("Failed on unmarshalling response")
	}

	expected := helloWorldResponse{
		Message: "hello world",
	}
	if responseJson != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", responseJson, expected)
	}

	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("content type header does not match: got %v want %v",
			ctype, "application/json")
	}
}
