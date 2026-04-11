// method-routing.go — Handling GET, POST, and DELETE on the same endpoint.
//
// Browsers send GET when you visit a URL, but real APIs use different HTTP
// methods for different operations (GET to read, POST to create, DELETE
// to remove). This file
// shows how one handler function can respond to multiple methods using
// r.Method and a switch statement. POST sends data in the request body;
// DELETE identifies the target via query parameters.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// In-memory data store — a simple map acting as our "database."
// This lives outside any function so all handlers can access it.
var services = map[string]string{
	"Google": "https://google.com",
	"GitHub": "https://github.com",
}

// servicesHandler routes requests based on HTTP method.
// One path ("/services") handles three different operations.
func servicesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		// GET /services — return all services as JSON.
		// Same pattern as json-response.go: set content type, encode, done.
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(services)

	case http.MethodPost:
		// POST /services — add a new service from the request body.
		// The client sends JSON: {"name": "Netflix", "url": "https://netflix.com"}
		// We decode it into a map, extract the fields, and add to our store.
		var newService map[string]string
		err := json.NewDecoder(r.Body).Decode(&newService)
		if err != nil {
			// Bad JSON — tell the client what went wrong.
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		name := newService["name"]
		url := newService["url"]
		if name == "" || url == "" {
			http.Error(w, "name and url required", http.StatusBadRequest)
			return
		}

		// Add to our store and respond with 201 Created.
		// 201 means "I created the resource you asked for."
		services[name] = url
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Added %s\n", name)

	case http.MethodDelete:
		// DELETE /services?name=Google — remove a service by name.
		// DELETE uses query parameters (in the URL) rather than a body.
		// r.URL.Query().Get("name") reads the ?name= value.
		name := r.URL.Query().Get("name")
		if name == "" {
			http.Error(w, "name parameter required", http.StatusBadRequest)
			return
		}

		// Check if the service exists before deleting — comma-ok pattern.
		if _, exists := services[name]; !exists {
			http.Error(w, "service not found", http.StatusNotFound)
			return
		}

		delete(services, name)
		fmt.Fprintf(w, "Deleted %s\n", name)

	default:
		// Any other method (PUT, PATCH, etc.) gets a 405.
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/services", servicesHandler)

	fmt.Println("Server starting on http://localhost:8080")
	fmt.Println("Try these curl commands in another terminal:")
	fmt.Println("  GET    curl http://localhost:8080/services")
	fmt.Println("  POST   curl -X POST -d '{\"name\":\"Netflix\",\"url\":\"https://netflix.com\"}' http://localhost:8080/services")
	fmt.Println("  DELETE curl -X DELETE 'http://localhost:8080/services?name=Google'")
	fmt.Println("Press Ctrl+C to stop")

	// log.Fatal prints an error and exits if the server can't start
	// (e.g., port 8080 is already in use by another program).
	log.Fatal(http.ListenAndServe(":8080", nil))
}
