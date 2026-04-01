package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Health Monitor v1")
}

// Return JSON using a map — quick and simple
func statusHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"service": "health-monitor",
		"status":  "running",
		"version": "1.0",
	}

	// Two lines to return JSON:
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Return a list of items as JSON
func servicesHandler(w http.ResponseWriter, r *http.Request) {
	services := []map[string]string{
		{"name": "Google", "url": "https://google.com"},
		{"name": "GitHub", "url": "https://github.com"},
		{"name": "Go.dev", "url": "https://go.dev"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/services", servicesHandler)

	fmt.Println("Server starting on http://localhost:8080")
	fmt.Println("Try these endpoints:")
	fmt.Println("  http://localhost:8080/")
	fmt.Println("  http://localhost:8080/status")
	fmt.Println("  http://localhost:8080/services")
	fmt.Println("Press Ctrl+C to stop")

	http.ListenAndServe(":8080", nil)
}
