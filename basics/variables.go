package main

import "fmt"

func main() {
	// Explicit type declaration
	var name string = "Alice"
	var age int = 25
	var price float64 = 19.99
	var isActive bool = true

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Price:", price)
	fmt.Println("Active:", isActive)

	// Short declaration — type inferred (most common)
	city := "Philadelphia"
	score := 95.5
	enrolled := true

	fmt.Println("\nCity:", city)
	fmt.Println("Score:", score)
	fmt.Println("Enrolled:", enrolled)

	// Multiple variables at once
	x, y := 10, 20
	fmt.Println("\nx:", x, "y:", y)

	// Constants
	const maxRetries = 3
	const appName = "Health Monitor"
	fmt.Println("\nApp:", appName, "Max retries:", maxRetries)

	// Try uncommenting this line — Go won't compile with unused variables
	// unused := "this will cause a compile error"
}
