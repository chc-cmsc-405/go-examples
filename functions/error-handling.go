package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Go's error pattern: return the result AND an error
// The caller checks the error before using the result

// Validate an age string — returns the age or an error
func parseAge(input string) (int, error) {
	age, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("invalid age: " + input)
	}
	if age < 0 || age > 150 {
		return 0, errors.New("age out of range: " + input)
	}
	return age, nil
}

// Check if a service URL is valid (simplified)
func validateURL(url string) error {
	if url == "" {
		return errors.New("URL cannot be empty")
	}
	if len(url) < 8 {
		return errors.New("URL too short: " + url)
	}
	return nil // nil means no error — success
}

func main() {
	// Pattern: call function, check error, then use result
	fmt.Println("=== Parsing Ages ===")

	inputs := []string{"25", "abc", "-5", "200", "30"}
	for _, input := range inputs {
		age, err := parseAge(input)
		if err != nil {
			fmt.Printf("  %q → Error: %s\n", input, err)
		} else {
			fmt.Printf("  %q → Age: %d\n", input, age)
		}
	}

	// Pattern: function that returns only an error
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

	// The if err != nil pattern — you'll write this hundreds of times
	// It's verbose, but every error is visible in the code
	// No hidden exceptions, no surprise crashes
}
