// channels-basic.go — A channel as a typed pipe between two goroutines.
//
// A channel synchronizes the sender and the receiver. An unbuffered channel
// (capacity 0, the default) makes the send and the receive a single event:
// the send only completes once a receiver is ready, and vice versa. The
// channel itself never holds the value. It hands it directly across.
//
// Compare to: Java BlockingQueue, Python queue.Queue. Channels are similar
// in spirit but have first-class language syntax (the <- operator) and an
// unbuffered default that encourages synchronous handoff.
//
// To run only this file:
//   go run channels-basic.go
package main

import "fmt"

func main() {
	// Create an unbuffered channel of strings.
	ch := make(chan string)

	// Spawn a goroutine that sends one value and exits.
	go func() {
		ch <- "hello from goroutine"
	}()

	// Receive blocks until the goroutine sends. The handoff is synchronous.
	msg := <-ch
	fmt.Println(msg)
}
