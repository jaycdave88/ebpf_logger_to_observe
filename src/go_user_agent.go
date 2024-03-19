package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type LogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
}

// Updated to match the provided curl command
func HTTPPost(endpoint string, token string, data []LogEntry) error {
	// Wrap the log entries in a map with a "data" key to match the curl example
	payload := map[string][]LogEntry{"data": data}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshaling data: %v", err)
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating POST request: %v", err)
	}

	// Use the token directly from the environment variable
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending POST request: %v", err)
	}
	defer resp.Body.Close()

	// Check the HTTP status code
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body) // Read the body for additional context in the error message
		return fmt.Errorf("received non-200 response status: %d - Body: %s", resp.StatusCode, string(body))
	}

	return nil
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// Get HTTP endpoint and bearer token from environment variables
	httpEndpoint := os.Getenv("OBSERVE_HTTP_ENDPOINT") // Ensure this matches your .env
	bearerToken := os.Getenv("BEARER_TOKEN")           // Ensure this matches your .env

	logs := []LogEntry{
		{Timestamp: time.Now(), Message: "Log entry 1"},
		{Timestamp: time.Now(), Message: "Log entry 2"},
	}

	err = HTTPPost(httpEndpoint, bearerToken, logs)
	if err != nil {
		fmt.Printf("Failed to send logs: %v\n", err)
		return
	}

	fmt.Println("Logs sent successfully")
}