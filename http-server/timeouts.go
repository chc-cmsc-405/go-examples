// timeouts.go — Setting timeouts on HTTP clients so requests don't hang.
//
// The default http.Get() has no timeout — if a server is slow or unreachable,
// your program waits indefinitely. In a health monitor, one slow URL could
// freeze your entire /check-all endpoint. The fix: create an http.Client
// with an explicit Timeout. After the timeout expires, the request fails
// with an error — your server stays responsive.
//
// This is a "good to know" concept for production code. The health monitor
// practice focuses on basic error handling (UP/DOWN), but timeouts are
// what separates demo code from real-world code.
package main

import (
	"fmt"
	"net/http"
	"time"
)

// checkWithTimeout wraps http.Get with a configurable timeout.
// Returns the status code, how long the request took, and any error.
func checkWithTimeout(url string, timeout time.Duration) (int, time.Duration, error) {
	// Create a NEW client with the timeout. Don't reuse http.DefaultClient
	// because changing its timeout would affect all requests globally.
	client := &http.Client{
		Timeout: timeout,
	}

	start := time.Now()
	resp, err := client.Get(url)
	elapsed := time.Since(start)

	if err != nil {
		// Could be a timeout, DNS failure, connection refused — we don't
		// distinguish here. The error message includes the reason.
		return 0, elapsed, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	return resp.StatusCode, elapsed, nil
}

func main() {
	// Three test URLs: one that works, one that's too slow, one that returns 500.
	urls := []string{
		"https://httpbin.org/get",       // responds quickly
		"https://httpbin.org/delay/3",   // delays 3 seconds — will exceed our 2s timeout
		"https://httpbin.org/status/500", // responds quickly but with an error status
	}

	timeout := 2 * time.Second
	fmt.Printf("Checking URLs with %v timeout:\n\n", timeout)

	for _, url := range urls {
		status, elapsed, err := checkWithTimeout(url, timeout)
		if err != nil {
			// The delay/3 URL should fail here — our 2s timeout fires before
			// the 3s delay finishes. The error message will include
			// "context deadline exceeded."
			fmt.Printf("  %-40s ERROR  (%v) — %v\n", url, elapsed.Round(time.Millisecond), err)
		} else {
			fmt.Printf("  %-40s %d    (%v)\n", url, status, elapsed.Round(time.Millisecond))
		}
	}
}
