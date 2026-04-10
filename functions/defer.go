// defer.go — Schedule cleanup to run when a function returns.
//
// In Python you use "with" statements or try/finally for cleanup.
// In Java you use try-with-resources. In Go, you use defer.
// defer guarantees the cleanup runs even if the function returns early
// due to an error. You'll use it every time you open an HTTP response
// body: defer resp.Body.Close(). Forgetting this causes resource leaks.
package main

import (
	"fmt"
	"net/http"
)

func main() {
	// === defer basics ===

	// defer pushes a function call onto a stack. It runs when main() returns,
	// no matter how or where main() exits. Watch the output order.
	fmt.Println("=== defer basics ===")
	fmt.Println("First")
	defer fmt.Println("This runs last (deferred)")
	fmt.Println("Second")
	fmt.Println("Third")
	// Output: First, Second, Third, then "This runs last" — after main() finishes

	// === Real-world use: closing HTTP response bodies ===

	// Every http.Get returns a response with a Body that must be closed.
	// Put defer right after the error check — this guarantees cleanup
	// even if later code in this function returns early.
	fmt.Println("\n=== defer with HTTP ===")
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Println("Error:", err)
		return // defer still runs even on early return
	}
	defer resp.Body.Close() // Scheduled now, runs when main() exits

	fmt.Println("Status:", resp.Status)
	fmt.Println("Content-Type:", resp.Header.Get("Content-Type"))

	// === Multiple defers: LIFO order ===

	// If you defer multiple calls, they run in reverse order (stack behavior).
	// Last deferred = first to run. This matters when you have multiple
	// resources to close — they close in reverse order of opening.
	fmt.Println("\n=== multiple defers ===")
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	defer fmt.Println("Deferred 3")
	fmt.Println("Before defers execute")
	// Output: "Before defers execute", then Deferred 3, 2, 1 (reverse order)
}
