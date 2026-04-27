// waitgroup-channel.go — WaitGroup + channel together.
//
// The most common Phase 4 pattern: spawn one goroutine per task, send each
// result on a channel, use a WaitGroup to know when every goroutine has
// finished, and close the channel so the receiver can stop ranging.
//
// Three pieces working together:
//   - WaitGroup counts how many goroutines are still running
//   - Channel carries each result back
//   - A separate goroutine waits on the WaitGroup and closes the channel,
//     which terminates the `for range ch` loop on the receiver.
//
// Without the closer goroutine, the range loop would block forever after
// the last receive, since the receiver cannot tell whether more values
// are coming.
//
// To run only this file:
//   go run waitgroup-channel.go
package main

import (
	"fmt"
	"sync"
	"time"
)

type result struct {
	ID    int
	Value string
}

func work(id int, out chan<- result) {
	time.Sleep(100 * time.Millisecond)
	out <- result{ID: id, Value: fmt.Sprintf("payload from %d", id)}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan result, 5)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			work(id, ch)
		}(i)
	}

	// Closer goroutine: wait for every worker to finish, then close the
	// channel so the range loop below can exit.
	go func() {
		wg.Wait()
		close(ch)
	}()

	for r := range ch {
		fmt.Printf("received id=%d value=%q\n", r.ID, r.Value)
	}
}
