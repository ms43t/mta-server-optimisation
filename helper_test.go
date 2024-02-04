package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

// TestGetInefficientInstance checks the getInefficientInstance function.

func TestGetInefficientInstance(t *testing.T) {
	threshold := 1
	expectedResult := []string{"mta-prod-1", "mta-prod-3"}

	result := getInefficientInstance(threshold)

	if len(result) != len(expectedResult) {
		t.Errorf("Expected %d inefficient instances, but got %d", len(expectedResult), len(result))
	}
}

// TestGetInstanceName simulates HTTP requests and verifies the responses.

func TestGetInstanceName(t *testing.T) {
	testCases := []struct {
		Name           string
		Request        *http.Request
		ExpectedResult []string
	}{
		{
			Name: "Valid Request",
			Request: httptest.NewRequest(http.MethodGet, "/mta-hosting-optimizer", nil).
				WithContext(setEnvContext("X", "1")),
			ExpectedResult: []string{"mta-prod-3"},
		},
		{
			Name: "Invalid Threshold",
			Request: httptest.NewRequest(http.MethodGet, "/mta-hosting-optimizer", nil).
				WithContext(setEnvContext("X", "invalid")),
			ExpectedResult: nil,
		},
		{
			Name: "Missing Threshold",
			Request: httptest.NewRequest(http.MethodGet, "/mta-hosting-optimizer", nil).
				WithContext(setEnvContext("", "")),
			ExpectedResult: []string{"mta-prod-1", "mta-prod-3"},
		},
		{
			Name: "Non-Get Request",
			Request: httptest.NewRequest(http.MethodPost, "/mta-hosting-optimizer", nil).
				WithContext(setEnvContext("X", "2")),
			ExpectedResult: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			w := httptest.NewRecorder()
			// Print debug information
			fmt.Printf("Environment variables: X=%s\n", tc.Request.Context().Value("X"))
			getInstanceName(w, tc.Request)

			if w.Code != http.StatusOK {
				t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
			}

			var response []string
			if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
				t.Errorf("Error decoding response body: %s", err)
			}
		})
	}

}

// TestGetEnv checks the getEnv function.

func TestGetEnv(t *testing.T) {
	testCases := []struct {
		Name           string
		Key            string
		DefaultValue   string
		ExpectedResult string
	}{
		{
			Name:           "Existing Key",
			Key:            "X",
			DefaultValue:   "1",
			ExpectedResult: "1",
		},
		{
			Name:           "Non-Existing Key",
			Key:            "Y",
			DefaultValue:   "2",
			ExpectedResult: "2",
		},
		{
			Name:           "Empty Key",
			Key:            "",
			DefaultValue:   "3",
			ExpectedResult: "3",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := getEnv(tc.Key, tc.DefaultValue)
			if result != tc.ExpectedResult {
				t.Errorf("getEnv returned wrong result for key '%s': got '%s', want '%s'", tc.Key, result, tc.ExpectedResult)
			}
		})
	}
}

// setEnvContext creates a context with environment variables.

func setEnvContext(key, value string) context.Context {
	env := make(map[string]string)
	env[key] = value
	return context.WithValue(context.Background(), "env", env)
}

// TestGoDotEnvVariable tests the GoDotEnvVariable function.

func TestGoDotEnvVariable(t *testing.T) {
	// Prepare a test .env file with sample key-value pairs
	envData := []byte(`X=1`)

	// Create a temporary .env file for testing
	err := os.WriteFile(".env", envData, 0644)
	if err != nil {
		t.Fatal("Failed to create .env file for testing:", err)
	}
	//defer func() {
	//	err := os.Remove(".env")
	//	if err != nil {
	//		log.Println("Failed to remove .env file after testing:", err)
	//	}
	//}()

	// Load the test .env file
	err = godotenv.Load(".env")
	if err != nil {
		t.Fatal("Failed to load .env file for testing:", err)
	}

	// Test cases
	testCases := []struct {
		Key           string
		ExpectedValue string
		DefaultValue  string
	}{
		{
			Key:           "KEY",
			ExpectedValue: "",
			DefaultValue:  "default_value",
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.Key, func(t *testing.T) {
			value := GoDotEnvVariable(tc.Key)

			if value != tc.ExpectedValue {
				t.Errorf("Unexpected value for key '%s'. Expected: '%s', Got: '%s'", tc.Key, tc.ExpectedValue, value)
			}
		})
	}
}
