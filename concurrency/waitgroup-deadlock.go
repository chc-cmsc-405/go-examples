// waitgroup-deadlock.go — What a missing wg.Done() looks like.
//
// If you call wg.Add(1) but the goroutine never calls wg.Done(), wg.Wait()
// blocks forever. The Go runtime detects this when every goroutine is asleep
// and exits with:
//
//   fatal error: all goroutines are asleep - deadlock!
//
// Recognize this message. It is the runtime telling you a WaitGroup is
// missing a Done somewhere.
//
// To run only this file:
//   go run waitgroup-deadlock.go
package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		// forgot wg.Done() — Wait() below will block forever
	}()
	wg.Wait()
}
