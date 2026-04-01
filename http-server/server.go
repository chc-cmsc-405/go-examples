package main

import (
	"fmt"
	"net/http"
)

// A handler function takes a ResponseWriter and a Request
// w = where you write the response (like System.out for the browser)
// r = the incoming request (URL, method, headers, body)
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go!")
}

func main() {
	// Map a URL path to a handler function
	http.HandleFunc("/", homeHandler)

	fmt.Println("Server starting on http://localhost:8080")
	fmt.Println("Press Ctrl+C to stop")

	// Start the server — this blocks (runs forever until you stop it)
	http.ListenAndServe(":8080", nil)
}
