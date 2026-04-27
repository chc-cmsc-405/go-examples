// waitgroup-basic.go — Wait until N goroutines finish.
//
// sync.WaitGroup is a counter:
//   wg.Add(n)  — raise the counter by n
//   wg.Done()  — lower it by 1
//   wg.Wait()  — block until the counter reaches 0
//
// Two ordering rules:
//   1. Call Add BEFORE the `go` statement, not inside the goroutine. Otherwise
//      Wait can return before the goroutine has even had a chance to add.
//   2. Use `defer wg.Done()` at the top of the goroutine, so Done still runs
//      if the function panics or returns early.
//
// Compare to: Java CountDownLatch, Python threading.Barrier. WaitGroup is
// closer to CountDownLatch, but with per-spawn Add(1) instead of a fixed
// count up front.
//
// To run only this file:
//   go run waitgroup-basic.go
package main

import (
	"fmt"
	"sync"
	"time"
)

func work(id int) {
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // 1. increment BEFORE spawning
		go func(id int) {
			defer wg.Done() // 2. defer guarantees Done runs
			work(id)
		}(i)
	}

	wg.Wait() // blocks until every Done() has been called
	fmt.Println("all workers finished")
}
