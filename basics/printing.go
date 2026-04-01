package main

import "fmt"

func main() {
	name := "Alice"
	age := 25
	score := 95.678
	active := true

	// Println — prints with newline
	fmt.Println("Hello,", name)

	// Printf — formatted output (like C's printf)
	fmt.Printf("Name: %s, Age: %d\n", name, age)
	fmt.Printf("Score: %.2f\n", score)  // 2 decimal places
	fmt.Printf("Active: %t\n", active)

	// %v — prints any value (your best friend for debugging)
	fmt.Printf("Value: %v\n", name)
	fmt.Printf("Value: %v\n", age)
	fmt.Printf("Value: %v\n", score)

	// Sprintf — returns a string instead of printing
	message := fmt.Sprintf("%s is %d years old", name, age)
	fmt.Println("\n" + message)

	// Common format verbs:
	// %s  string
	// %d  integer
	// %f  float (%.2f for 2 decimals)
	// %t  boolean
	// %v  any value (default format)
	// %+v struct with field names
	// %q  quoted string
}
