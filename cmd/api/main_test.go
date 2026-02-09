package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T) {
	tests := []struct {
		name           string
		handler        http.HandlerFunc
		expectedStatus int
		expectedBody   string
	}{
		{"Home", home, http.StatusOK, "Welcome to the Shapes API"},
		{"Health", health, http.StatusOK, "Server is running"},
		{"About", about, http.StatusOK, "Alfred Acosta"},
		{"Greeting", greeting, http.StatusOK, "Hello from the Shapes API"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			rr := httptest.NewRecorder()

			tt.handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("expected %v, got %v", tt.expectedStatus, rr.Code)
			}

			if rr.Body.String() != tt.expectedBody {
				t.Errorf("expected %v, got %v", tt.expectedBody, rr.Body.String())
			}
		})
	}
}

