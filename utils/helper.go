package utils

import (
	"strings"
	"fmt"
	"log"
	"net/http"
	"io"
	"encoding/json"
)

// GetServiceName returns the service name from the pod name
func getServiceName(podName string) string {
	parts := strings.Split(podName, "-")
	return strings.Join(parts[:len(parts)-2], "-")
}

// SendResponse sends a message to both the client and the console
func SendResponse(w http.ResponseWriter, message string) {
	fmt.Fprintln(w, message)
	log.Println(message)
}

// Helper function to validate request and read body
func validateAndReadBody(w http.ResponseWriter, r *http.Request) ([]byte, error) {
    if r.Method != "POST" {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return nil, fmt.Errorf("invalid request method")
    }

    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Failed to read request body", http.StatusInternalServerError)
        return nil, fmt.Errorf("failed to read request body: %v", err)
    }

    return body, nil
}

// Helper function to debug and unmarshal JSON payload
func debugAndUnmarshal(body []byte, payload interface{}, handlerName string) error {
    if AppConfig.EnableDebug {
        log.Printf("Payload %s: \n%s", handlerName, body)
    }
    err := json.Unmarshal(body, payload)
    if err != nil {
        return fmt.Errorf("invalid JSON format: %v", err)
    }
    return nil
}