package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWebHealthcheck(t *testing.T) {
	tests := []struct {
		name     string
		expected int
	}{
		{
			name:     "success",
			expected: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/web-healthcheck", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(WebHealthcheck)

			ctx := req.Context()
			req = req.WithContext(ctx)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expected {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expected)
			}
		})
	}
}
