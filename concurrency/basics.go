// basics.go — Spawning concurrent work with the `go` keyword.
//
// A goroutine is the smallest unit of concurrent work in Go. The `go` keyword
// in front of a function call runs that function on its own goroutine and
// returns control to the caller immediately. The function itself runs in the
// background, scheduled by the Go runtime onto a small pool of OS threads.
//
// A goroutine costs roughly what a function call costs (~2 KB initial stack).
// Spawning thousands is reasonable. Spawning thousands of OS threads is not.
//
// Compare to: Java `new Thread(...).start()`, Python `threading.Thread(...).start()`.
// Both are real OS threads with ~1-8 MB stacks. Go's runtime-managed goroutines
// are much cheaper, which is why "spawn one per task" is the idiomatic pattern.
package main

import (
	"fmt"
	"time"
)

// === The work each goroutine performs ===

// `work` simulates a slow task (e.g., an HTTP request). The 100 ms sleep stands
// in for network latency. The id parameter lets us see which goroutine printed
// each line.
func work(id int) {
	fmt.Printf("goroutine %d starting\n", id)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("goroutine %d done\n", id)
}

func main() {
	// === Sequential — for comparison ===
	// Without `go`, each call blocks until it returns. Three calls take 300 ms.
	fmt.Println("--- sequential ---")
	start := time.Now()
	for i := 1; i <= 3; i++ {
		work(i)
	}
	fmt.Println("sequential elapsed:", time.Since(start))

	// === Concurrent — three goroutines ===
	// Each `go work(i)` starts the function on its own goroutine and returns
	// immediately. All three run in parallel.
	fmt.Println("\n--- concurrent ---")
	start = time.Now()
	for i := 1; i <= 3; i++ {
		go work(i)
	}

	// Crude wait. The program will exit before the goroutines finish if main
	// returns. A real program uses sync.WaitGroup (see waitgroups example).
	time.Sleep(150 * time.Millisecond)
	fmt.Println("concurrent elapsed:", time.Since(start))

	// Notice the output:
	// - Sequential: each "starting" is followed by its own "done"
	// - Concurrent: all three "starting" lines appear, then all three "done"
	//   lines, often in different orders across runs
}
