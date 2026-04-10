// functions.go — Declaring functions, parameters, return values, and control flow.
//
// Go puts types after parameter names: func add(a, b int) int. This reads
// left to right as "add takes a and b, both int, returns int." Go also has
// only ONE loop keyword: for. No while, no do-while. The for loop covers
// all three patterns depending on how you write it. And if/else has no
// parentheses around the condition — braces are required instead.
package main

import "fmt"

// === Functions ===

// Basic function — types come AFTER parameter names (not before like Java).
// When consecutive parameters share a type, you can group them: (a, b int).
func add(a, b int) int {
	return a + b
}

// String parameter, string return.
func greet(name string) string {
	return "Hello, " + name + "!"
}

// Multiple parameters with different types.
// Sprintf returns a formatted string — like Python's f-string.
func describe(name string, age int) string {
	return fmt.Sprintf("%s is %d years old", name, age)
}

// === Control flow: if/else ===

// No parentheses around the condition — Go uses braces instead.
// Compare to Python (indentation-based) and Java (parentheses required).
func checkAge(age int) string {
	if age >= 21 {
		return "Full access"
	} else if age >= 18 {
		return "Limited access"
	} else {
		return "No access"
	}
}

// === Loops: Go only has "for" ===

// Traditional for loop — identical to C/Java.
func sum(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}
	return total
}

// "While" style — just for with a condition, no init or post statement.
// This is how Go replaces while loops.
func countdown(from int) {
	for from > 0 {
		fmt.Println(from)
		from--
	}
	fmt.Println("Go!")
}

// Note: for {} with no condition is an infinite loop (like while True in Python).

func main() {
	fmt.Println("add(3, 5):", add(3, 5))
	fmt.Println(greet("Alice"))
	fmt.Println(describe("Bob", 30))
	fmt.Println("checkAge(25):", checkAge(25))
	fmt.Println("checkAge(18):", checkAge(18))
	fmt.Println("checkAge(15):", checkAge(15))
	fmt.Println("sum(10):", sum(10))
	fmt.Println("\nCountdown:")
	countdown(3)
}
