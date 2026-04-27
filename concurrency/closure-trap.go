// closure-trap.go — The classic loop-variable capture bug, and the safe pattern.
//
// On Go versions before 1.22, this loop is a common bug:
//
//   for i := 0; i < 5; i++ {
//       go func() { fmt.Println(i) }()
//   }
//
// Every closure captures the same variable `i`, so all five goroutines often
// print the same final value (5) instead of 0–4. Go 1.22 changed loop semantics
// so each iteration creates a fresh `i`, fixing the bug for new code.
//
// The course is on Go 1.22+, but the explicit "pass the value as an argument"
// pattern still appears in published code, so recognize it when you see it.
//
// To run only this file:
//   go run closure-trap.go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// === Version A — captures the loop variable ===
	// On Go 1.22+ each iteration gets a fresh `i`, so this prints 0–4 in some
	// order. On older Go, all five goroutines saw the same `i` and typically
	// printed `5` five times.
	fmt.Println("--- version A: closure captures i ---")
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("A:", i)
		}()
	}
	wg.Wait()

	// === Version B — passes i as an argument ===
	// Each goroutine gets its own copy of the value via the parameter `n`.
	// This pattern works on every Go version, so it is the safer default
	// when writing portable code.
	fmt.Println("\n--- version B: pass i as an argument ---")
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println("B:", n)
		}(i)
	}
	wg.Wait()

	// Both versions print 0–4 in some order on Go 1.22+. The difference shows
	// up if you run on Go 1.21 or earlier, or in any code that captures a
	// variable from outside a `for` loop (e.g., from a slice index or a
	// channel receive). Passing the value explicitly is the rule that always
	// works.
}
