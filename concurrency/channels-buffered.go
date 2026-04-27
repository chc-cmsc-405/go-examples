// channels-buffered.go — A buffered channel decouples sender and receiver.
//
// A buffered channel can hold N values before the next send blocks. This is
// useful when you know exactly how many values will be produced and want
// every producer to send and exit without waiting. The receiver collects at
// its own pace.
//
// The pattern below is the natural fit for Phase 4's concurrent /check-all:
// one buffered channel sized to len(services) lets every goroutine send its
// result and exit. The handler then reads len(services) values out.
//
// To run only this file:
//   go run channels-buffered.go
package main

import (
	"fmt"
	"time"
)

// worker simulates a slow task that produces a result string.
// `chan<- string` declares an "out" channel: this function can only send,
// not receive. The compiler enforces the direction.
func worker(id int, out chan<- string) {
	time.Sleep(100 * time.Millisecond)
	out <- fmt.Sprintf("worker %d done", id)
}

func main() {
	const N = 5

	// Buffer sized to N, so every worker can send without waiting.
	ch := make(chan string, N)

	for i := 1; i <= N; i++ {
		go worker(i, ch)
	}

	// Counting receives — same number as sends, no WaitGroup needed.
	for i := 0; i < N; i++ {
		fmt.Println(<-ch)
	}
}
