package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestGetAlbumByID(t *testing.T) {
	// Get the service URL from an environment variable (API_URL)
	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		t.Fatalf("No environment variable set for API_URL.")
	}

	// Make an actual HTTP request to the API
	resp, err := http.Get(apiURL)
	if err != nil {
		t.Fatalf("Failed to make the request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read the response body: %v", err)
	}

	expected := `{"id":"1","title":"Blue Train","artist":"John Coltrane","price":56.99}`

	// Check the HTTP response code and body
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.JSONEq(t, expected, string(body))

	fmt.Println("Test passed: Get Album By ID")
}
