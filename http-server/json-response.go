// json-response.go — Encoding Go data as JSON and sending it in HTTP responses.
//
// Most APIs return JSON, not plain text. Go's encoding/json package converts
// maps and structs to JSON automatically. The pattern is two lines:
//   w.Header().Set("Content-Type", "application/json")
//   json.NewEncoder(w).Encode(data)
// The encoder writes directly to the ResponseWriter — no intermediate string
// needed. Compare to Python's json.dumps() which creates a string first.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Health Monitor v1")
}

// === Returning a JSON object ===

// Use a map[string]string to build a JSON object on the fly.
// The map keys become JSON keys; the values become JSON values.
func statusHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"service": "health-monitor",
		"status":  "running",
		"version": "1.0",
	}

	// Set the content type BEFORE writing — tells the client to expect JSON.
	// json.NewEncoder(w) writes directly to the HTTP response stream.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// === Returning a JSON array ===

// A slice of maps becomes a JSON array of objects.
// Each map in the slice becomes one object in the array.
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

	// log.Fatal prints an error and exits if the server can't start
	// (e.g., port 8080 is already in use by another program).
	log.Fatal(http.ListenAndServe(":8080", nil))
}
