package auth

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	expected := "asd8dfadfa8dasfads8"
	req := httptest.NewRequest(http.MethodGet, "/v1/users", nil)
	req.Header.Set("Authorization", fmt.Sprintf("ApiKey %v", expected))

	apiKey, err := GetAPIKey(req.Header)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if apiKey != expected {
		t.Errorf("Expected %v but got %v", expected, apiKey)
	}
}