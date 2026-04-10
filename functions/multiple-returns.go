// multiple-returns.go — Functions that return two (or more) values.
//
// This is one of Go's most distinctive features. Most languages return one
// value; Go routinely returns two. The most common pattern: (result, error).
// The caller is expected to check the error BEFORE using the result. This
// replaces try/catch entirely — errors are values, not exceptions.
//
// Python can return tuples (result, error) but rarely does by convention.
// Java uses exceptions. Go makes the two-value return the standard pattern.
package main

import (
	"errors"
	"fmt"
)

// === The (result, error) pattern ===

// divide returns the result AND an error. If b is zero, the result is
// meaningless — the caller must check err before using the float.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil // nil = no error, everything is fine
}

// === Other multiple return patterns ===

// Swap returns two strings in reversed order.
// Multiple returns aren't just for errors — they're for any case
// where a function naturally produces two values.
func swap(a, b string) (string, string) {
	return b, a
}

// findUser returns a value and a boolean indicating whether it was found.
// This is the "comma ok" pattern — you'll see it with map lookups too.
func findUser(id int) (string, bool) {
	users := map[int]string{
		1: "Alice",
		2: "Bob",
		3: "Charlie",
	}
	name, found := users[id]
	return name, found
}

func main() {
	// === Using (result, error) ===

	// Always check err before using result. This pattern becomes muscle memory.
	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 3 = %.2f\n", result)
	}

	// Divide by zero — triggers the error path.
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 0 = %.2f\n", result)
	}

	// === Swap ===
	first, second := swap("hello", "world")
	fmt.Printf("\nSwap: %s, %s\n", first, second)

	// === Comma-ok pattern ===

	// Same shape as (result, error) but the second value is a bool.
	// You'll see this with map lookups: value, ok := myMap[key]
	name, found := findUser(2)
	if found {
		fmt.Printf("\nFound user: %s\n", name)
	}

	name, found = findUser(99)
	if !found {
		fmt.Println("User 99 not found")
	}

	// === The blank identifier _ ===

	// Use _ when you need to ignore a return value.
	// Go won't let you declare a variable you don't use — _ is the escape hatch.
	_, err = divide(5, 0)
	fmt.Println("\nIgnored result, got error:", err)
}
