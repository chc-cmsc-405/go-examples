// http-client.go — Making outbound HTTP requests from Go with http.Get.
//
// Your server doesn't just receive requests — it makes them too. The health
// monitor calls other services to check if they're up. http.Get makes a
// GET request and returns a response + error. The pattern is always:
// call, check error, defer close body, use the response.
//
// This is where Go's error handling really matters. Network requests fail
// for many reasons (DNS, timeout, connection refused). Each failure
// returns an error — your code decides what "DOWN" means.
package main

import (
	"fmt"
	"net/http"
)

// checkURL makes an HTTP GET request and returns "UP" or "DOWN".
// This is the core pattern your health monitor will use.
func checkURL(url string) string {
	// http.Get returns (response, error). The error is non-nil if the
	// request failed entirely — DNS failure, connection refused, timeout.
	resp, err := http.Get(url)
	if err != nil {
		// Request failed — the site is unreachable.
		return "DOWN"
	}
	// Always close the response body when you're done.
	// defer ensures this runs even if the function returns early later.
	defer resp.Body.Close()

	// A 200 status code means the server responded successfully.
	// Anything else (404, 500, etc.) means something is wrong.
	if resp.StatusCode == 200 {
		return "UP"
	}
	return "DOWN"
}

func main() {
	// Test with real URLs — including one that will fail.
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://go.dev",
		"http://this-does-not-exist-abc123.com",
	}

	// Loop and check each URL. The %-45s pads the URL to 45 characters
	// so the status column lines up neatly.
	for _, url := range urls {
		status := checkURL(url)
		fmt.Printf("%-45s %s\n", url, status)
	}
}
