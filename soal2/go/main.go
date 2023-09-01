package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

// Response represents the JSON response structure.
type Response struct {
	Hostname string    `json:"hostname"`
	Time     time.Time `json:"time"`
	Message  string    `json:"message"`
}

// handler handles incoming HTTP requests.
func handler(w http.ResponseWriter, r *http.Request) {
	// Get the hostname of the server
	hostname, _ := os.Hostname()

	// Determine the welcome message based on the request's host
	var message string
	if r.Host == "domain.tbd" {
		message = "Welcome to the domain!"
	} else {
		message = "Welcome to the IP address!"
	}

	// Create a Response instance with hostname, current time, and message
	response := Response{
		Hostname: hostname,
		Time:     time.Now(),
		Message:  message,
	}

	// Marshal the Response to JSON
	jsonResponse, _ := json.Marshal(response)

	// Set response headers and write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func main() {
	// Handle the root path ("/") with the handler function
	http.Handle("/", http.HandlerFunc(handler))

	// Listen and serve on port 8080
	http.ListenAndServe(":8080", nil)
}
