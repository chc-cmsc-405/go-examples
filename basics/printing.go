// printing.go — Println, Printf, and Sprintf for formatted output.
//
// Go's fmt package handles all formatted output. Println is like Python's
// print() — simple, adds a newline. Printf is like C's printf — you control
// the format with verbs (%s, %d, %f). The most useful verb is %v, which
// prints any value in its default format — great for debugging when you
// don't want to think about types.
package main

import "fmt"

func main() {
	name := "Alice"
	age := 25
	score := 95.678
	active := true

	// === Println — simple output ===

	// Prints values separated by spaces, adds a newline.
	// Like Python's print("Hello,", name).
	fmt.Println("Hello,", name)

	// === Printf — formatted output ===

	// Format verbs control how values are displayed.
	// You must add \n yourself — Printf doesn't auto-newline.
	fmt.Printf("Name: %s, Age: %d\n", name, age)
	fmt.Printf("Score: %.2f\n", score) // .2f = 2 decimal places
	fmt.Printf("Active: %t\n", active)

	// === %v — the universal verb ===

	// %v prints any value in its default format. Use it when you don't
	// care about precise formatting — especially for debugging.
	fmt.Printf("Value: %v\n", name)
	fmt.Printf("Value: %v\n", age)
	fmt.Printf("Value: %v\n", score)

	// === Sprintf — format into a string ===

	// Like Printf but returns the string instead of printing it.
	// Useful when you need to build a string for later use.
	// Compare to Python's f-strings: f"{name} is {age} years old"
	message := fmt.Sprintf("%s is %d years old", name, age)
	fmt.Println("\n" + message)

	// === Quick reference ===
	//
	// %s  string          "Alice"
	// %d  integer         42
	// %f  float           19.990000
	// %.2f float (2 dec)  19.99
	// %t  boolean         true
	// %v  any value       (default format — works for everything)
	// %+v struct + fields {Name:Alice Age:25}
	// %q  quoted string   "Alice"
}
