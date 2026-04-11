// server.go — A minimal HTTP server in Go.
//
// Go's standard library includes a production-grade HTTP server. No
// frameworks, no dependencies, no configuration files. Two lines of code:
// register a handler function for a URL path, then start listening.
// Compare to Java (Spring Boot + annotations + config) or Python
// (Flask/Django + setup). Go's net/http package does it all directly.
//
// Note: log.Fatal wraps ListenAndServe so you see an error if the server
// can't start (e.g., port 8080 is already in use by another program).
// Without it, the server would silently fail to start.
package main

import (
	"fmt"
	"log"
	"net/http"
)

// === Handler functions ===

// Every HTTP handler has the same signature: (ResponseWriter, *Request).
// w = where you write the response (think of it as System.out for the browser).
// r = the incoming request — contains the URL, method, headers, and body.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Fprintf writes formatted text to any writer — here, the HTTP response.
	fmt.Fprintf(w, "Hello from Go!")
}

func main() {
	// === Registering routes ===

	// HandleFunc maps a URL path to a handler function.
	// When someone visits http://localhost:8080/, Go calls homeHandler.
	http.HandleFunc("/", homeHandler)

	fmt.Println("Server starting on http://localhost:8080")
	fmt.Println("Press Ctrl+C to stop")

	// === Starting the server ===

	// ListenAndServe blocks — the program runs forever, handling requests,
	// until you stop it. The ":8080" means "listen on port 8080 on all
	// network interfaces." The nil means "use the default request router."
	log.Fatal(http.ListenAndServe(":8080", nil))
}
