package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "ApiKey 123456")

	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if apiKey != "123456" {
		t.Errorf("Expected '123456', got '%s'", apiKey)
	}

	headers = http.Header{}
	headers.Add("Authorization", "Bearer 123456")

	_, err = GetAPIKey(headers)
	if err == nil {
		t.Errorf("Expected error for malformed authorization header, got nil")
	}
}
