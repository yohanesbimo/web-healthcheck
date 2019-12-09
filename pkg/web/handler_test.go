package web

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TestRenderHTML(t *testing.T) {
	tests := []struct {
		name     string
		template string
		expected error
	}{
		{
			name:     "success",
			template: "../../web/app/home.html",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpl, _ := template.ParseFiles(tt.template)
			rr := httptest.NewRecorder()
			result := renderHTML(rr, tmpl)
			if result != tt.expected {
				t.Error("failed check")
			}
		})
	}
}

func TestHomeView(t *testing.T) {
	tests := []struct {
		name     string
		template string
		expected int
	}{
		{
			name:     "success",
			expected: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(HomeView)

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
