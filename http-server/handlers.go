package main

import (
	"fmt"
	"net/http"
)

// Each handler is a function with the same signature
// Adding a new endpoint = one HandleFunc + one function

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
