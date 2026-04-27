// channels-range.go — Iterating over a channel until the sender closes it.
//
// `for v := range ch` keeps receiving until the channel is closed and drained.
// Closing is a one-way signal from the sender saying "no more values are
// coming." Without it, `for range` would block forever after the last
// receive, waiting for a value that never comes.
//
// Two rules go with close:
//   1. Only the sender closes — receivers that close are a bug, since other
//      senders would then panic on send.
//   2. Closing is final — calling close twice panics.
//
// To run only this file:
//   go run channels-range.go
package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	// Producer goroutine sends three values, then closes the channel.
	// `defer close(ch)` guarantees close runs even if the function returns
	// early or panics.
	go func() {
		defer close(ch)
		for i := 1; i <= 3; i++ {
			ch <- i
		}
	}()

	// Consumer ranges until the channel is closed and drained.
	for v := range ch {
		fmt.Println("got", v)
	}

	fmt.Println("range exited cleanly")
}
