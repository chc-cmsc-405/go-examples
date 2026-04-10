// error-handling.go — Go's explicit error handling: if err != nil.
//
// Go has no try/catch. When a function can fail, it returns an error value.
// The caller checks the error before using the result. This is the most
// distinctive pattern in Go — you'll write "if err != nil" hundreds of times.
//
// This file covers two levels:
//   1. Basic errors — creating errors with errors.New, returning them
//   2. Error wrapping — adding context with fmt.Errorf and %w so you know
//      WHERE in the call chain something failed, not just WHAT failed
//
// Compare to Python: try/except catches errors far from where they're thrown.
// In Go, every error is handled right where it occurs. More verbose, but
// you always know exactly what happens when something fails.
package main

import (
	"errors"
	"fmt"
	"strconv"
)

// === Basic errors: returning (value, error) ===

// parseAge converts a string to an int and validates the range.
// Returns (result, error) — the standard Go pattern. The caller
// must check err before using the age.
func parseAge(input string) (int, error) {
	// strconv.Atoi is Go's equivalent of Python's int() or Java's Integer.parseInt().
	// It returns (value, error) — the same pattern we're building here.
	age, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("invalid age: " + input)
	}
	if age < 0 || age > 150 {
		return 0, errors.New("age out of range: " + input)
	}
	return age, nil // nil = no error, everything is fine
}

// validateURL returns only an error (no result value).
// Some functions exist just to say "is this input valid?"
// nil means valid; non-nil means invalid with a reason.
func validateURL(url string) error {
	if url == "" {
		return errors.New("URL cannot be empty")
	}
	if len(url) < 8 {
		return errors.New("URL too short: " + url)
	}
	return nil
}

// === Error wrapping: adding context ===

// divide is a simple function that can fail (division by zero).
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// calculateAverage calls divide and wraps any error with context.
// fmt.Errorf("context: %w", err) creates a new error that includes
// the original error inside it. The %w verb preserves the error chain.
//
// Without wrapping: "cannot divide by zero"
// With wrapping:    "calculating average: cannot divide by zero"
// The second version tells you WHERE the failure happened.
func calculateAverage(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, errors.New("no values provided")
	}

	sum := 0.0
	for _, v := range values {
		sum += v
	}

	result, err := divide(sum, float64(len(values)))
	if err != nil {
		// Wrap the error with context — "calculating average:" is the WHERE.
		return 0, fmt.Errorf("calculating average: %w", err)
	}
	return result, nil
}

func main() {
	// === Using (value, error) — basic pattern ===

	fmt.Println("=== Parsing Ages ===")

	inputs := []string{"25", "abc", "-5", "200", "30"}
	for _, input := range inputs {
		age, err := parseAge(input)
		if err != nil {
			// Error path — print the error, don't use the result.
			fmt.Printf("  %q → Error: %s\n", input, err)
		} else {
			// Success path — safe to use the result.
			fmt.Printf("  %q → Age: %d\n", input, age)
		}
	}

	// === Using error-only returns ===

	fmt.Println("\n=== Validating URLs ===")

	urls := []string{"https://google.com", "", "short", "https://github.com"}
	for _, url := range urls {
		err := validateURL(url)
		if err != nil {
			fmt.Printf("  %q → Error: %s\n", url, err)
		} else {
			fmt.Printf("  %q → Valid\n", url)
		}
	}

	// The if err != nil pattern — you'll write this hundreds of times.
	// It's verbose, but every error is visible in the code.
	// No hidden exceptions, no surprise crashes.

	// === Using error wrapping ===

	fmt.Println("\n=== Error Wrapping ===")

	// Happy path — calculateAverage succeeds.
	avg, err := calculateAverage([]float64{120, 85, 200, 95})
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Average: %.1f\n", avg)
	}

	// Error path — empty input triggers the error.
	// The error message will say "no values provided" — from calculateAverage.
	_, err = calculateAverage([]float64{})
	if err != nil {
		fmt.Println("Error:", err)
	}
}
