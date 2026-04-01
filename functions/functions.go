package main

import "fmt"

// Basic function — types come after parameter names
func add(a, b int) int {
	return a + b
}

// String function
func greet(name string) string {
	return "Hello, " + name + "!"
}

// Multiple parameters with different types
func describe(name string, age int) string {
	return fmt.Sprintf("%s is %d years old", name, age)
}

// If/else — no parentheses around condition
func checkAge(age int) string {
	if age >= 21 {
		return "Full access"
	} else if age >= 18 {
		return "Limited access"
	} else {
		return "No access"
	}
}

// For loop — Go's ONLY loop (no while, no do-while)
func sum(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}
	return total
}

// "While" style loop using for
func countdown(from int) {
	for from > 0 {
		fmt.Println(from)
		from--
	}
	fmt.Println("Go!")
}

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
