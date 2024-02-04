package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// HostNameIPStatus represents the IP and status of a host.
type HostNameIPStatus struct {
	IP     string
	Status bool
}

// GetHostNameResponse represents the response structure for the GetInstanceName API.
type GetHostNameResponse struct {
	Result []string `json:"result"`
	Status string   `json:"status"`
	Error  error    `json:"error"`
}

// ipMap stores the mapping of host names to IP status.
var ipMap map[string][]HostNameIPStatus

// getInstanceName handles requests to the GetInstanceName API.
func getInstanceName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	threshold, err := strconv.Atoi(GoDotEnvVariable("X")) // Get threshold from environment.
	if err != nil {
		log.Println("Error converting string to int")
		json.NewEncoder(w).Encode(GetHostNameResponse{
			Result: nil,
			Status: "Error",
			Error:  err,
		})
		return
	}

	result := getInefficientInstance(threshold) // Get inefficient instances.
	json.NewEncoder(w).Encode(result)           // Encode and send the response.
}

// getInefficientInstance identifies inefficient instances based on the threshold.
func getInefficientInstance(threshold int) []string {
	inefficientInstance := make([]string, 0)

	for key, val := range ipMap {
		count := 0
		for _, ipStatus := range val {
			if ipStatus.Status {
				count++
			}
		}
		if count <= threshold {
			inefficientInstance = append(inefficientInstance, key)
		}
	}

	return inefficientInstance
}

// loadMockData initializes mock data for testing.
func loadMockData() error {
	ips := []string{"127.0.0.1", "127.0.0.2", "127.0.0.3", "127.0.0.4", "127.0.0.5", "127.0.0.6"}
	hostNames := []string{"mta-prod-1", "mta-prod-1", "mta-prod-2", "mta-prod-2", "mta-prod-2", "mta-prod-3"}
	actives := []bool{true, false, true, true, false, false}

	ipMap = make(map[string][]HostNameIPStatus)

	for idx := 0; idx < len(ips); idx++ {
		ipMap[hostNames[idx]] = append(ipMap[hostNames[idx]], NewHostNameIPStatus(ips[idx], actives[idx]))
	}

	return nil
}

// NewHostNameIPStatus is a constructor function for HostNameIPStatus.
func NewHostNameIPStatus(ip string, status bool) HostNameIPStatus {
	return HostNameIPStatus{IP: ip, Status: status}
}

// getEnv retrieves the value of an environment variable or returns a default value.
func getEnv(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

// GoDotEnvVariable retrieves the value of an environment variable using the "godotenv" package.
func GoDotEnvVariable(key string) string {
	err := godotenv.Load() // Load environment variables.
	if err != nil {
		log.Fatalf("Error loading environment variables")
	}

	return os.Getenv(key)
}
