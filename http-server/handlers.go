// handlers.go — Adding multiple endpoints to an HTTP server.
//
// Each endpoint is one HandleFunc call + one handler function. Adding a new
// page to your server is two lines of code: register the path, write the
// function. This pattern scales to any number of endpoints — your health
// monitor will have 10+ endpoints all registered the same way.
package main

import (
	"fmt"
	"net/http"
)

// === Handler functions ===

// Each handler has the same signature. The function name doesn't matter
// to Go — what matters is the path you register it with below.

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Health Monitor")
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status: OK")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Health Monitor v1.0 — built with Go")
}

func main() {
	// === Route registration ===

	// Each HandleFunc maps a URL path → handler function.
	// "/" matches the root; "/status" and "/about" are specific paths.
	// Order doesn't matter — Go's router matches the most specific path.
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/about", aboutHandler)

	fmt.Println("Server starting on http://localhost:8080")
	fmt.Println("Try these endpoints:")
	fmt.Println("  http://localhost:8080/")
	fmt.Println("  http://localhost:8080/status")
	fmt.Println("  http://localhost:8080/about")
	fmt.Println("Press Ctrl+C to stop")

	http.ListenAndServe(":8080", nil)
}
