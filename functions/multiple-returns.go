package main

import (
	"errors"
	"fmt"
)

// Multiple return values — Go's signature pattern
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Another example — swap two values
func swap(a, b string) (string, string) {
	return b, a
}

// Return a value and a boolean (common "found" pattern)
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
	// Divide — must handle both return values
	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 3 = %.2f\n", result)
	}

	// Divide by zero — triggers the error
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 0 = %.2f\n", result)
	}

	// Swap
	first, second := swap("hello", "world")
	fmt.Printf("\nSwap: %s, %s\n", first, second)

	// Find user — the "comma ok" pattern
	name, found := findUser(2)
	if found {
		fmt.Printf("\nFound user: %s\n", name)
	}

	name, found = findUser(99)
	if !found {
		fmt.Println("User 99 not found")
	}

	// Ignore a return value with _ (blank identifier)
	_, err = divide(5, 0)
	fmt.Println("\nIgnored result, got error:", err)
}
