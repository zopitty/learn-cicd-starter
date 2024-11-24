package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		headers       http.Header
		expectedKey   string
		expectedError error
	}{
		"valid api key": {
			headers: http.Header{
				"Authorization": []string{"ApiKey 12345"},
			},
			expectedKey:   "12345",
			expectedError: nil,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			key, err := GetAPIKey(tc.headers)
			if key != tc.expectedKey {
				t.Errorf("expected key: %s, got: %s", tc.expectedKey, key)
			}
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("expected error: %v, got %v", tc.expectedError, err)
			}
		})
	}
}
