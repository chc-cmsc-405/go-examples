// timing.go — Side-by-side sequential vs concurrent timing.
//
// The point of concurrency is observable in the timing difference. Sequential
// code that does I/O leaves the CPU idle while the network responds. Concurrent
// code starts every request before any of them returns, so all the waits
// happen in parallel.
//
// To run only this file (not basics.go), use:
//   go run timing.go
package main

import (
	"fmt"
	"time"
)

// fakeNetworkCall stands in for an HTTP request. The 200 ms sleep is the
// "waiting on the network" portion that concurrency lets us overlap.
func fakeNetworkCall(id int) {
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("call %d returned\n", id)
}

func main() {
	const N = 5

	// === Sequential ===
	// Each call blocks for 200 ms. Total: ~1 second.
	fmt.Println("--- sequential ---")
	start := time.Now()
	for i := 1; i <= N; i++ {
		fakeNetworkCall(i)
	}
	fmt.Printf("sequential: %v (%d calls × 200 ms)\n\n", time.Since(start), N)

	// === Concurrent ===
	// All N calls start within microseconds of each other. They all sleep
	// in parallel, so the total wall time is ~200 ms regardless of N.
	fmt.Println("--- concurrent ---")
	start = time.Now()
	for i := 1; i <= N; i++ {
		go fakeNetworkCall(i)
	}

	// Wait long enough for the goroutines to finish. A real program uses
	// sync.WaitGroup. For this demo, a sleep slightly longer than 200 ms is
	// enough.
	time.Sleep(250 * time.Millisecond)
	fmt.Printf("concurrent (with sleep): %v\n", time.Since(start))

	// Try N = 50 or N = 500. Sequential time scales linearly. Concurrent
	// time stays flat. That flat line is what makes Go feel different for
	// network-bound workloads.
}
