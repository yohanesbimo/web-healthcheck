package api

import (
	"testing"
)

func TestHealthchecker(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		expected bool
	}{
		{
			name:     "success",
			url:      "http://spacestock.com",
			expected: true,
		},
		{
			name:     "success",
			url:      "http://tokopedia.com",
			expected: false,
		},
		{
			name:     "success",
			url:      "http://abc.xyz",
			expected: true,
		},
		{
			name:     "failed",
			url:      "http://123.abc",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Healthchecker(tt.url)
			if result != tt.expected {
				t.Error("failed check")
			}
		})
	}
}
