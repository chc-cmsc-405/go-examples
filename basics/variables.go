// variables.go — Declaring variables with var and :=.
//
// Go has two ways to declare variables. var is explicit (you name the type),
// := is shorthand (the type is inferred from the value). You'll use := most
// of the time — it's shorter and the compiler figures out the type. The key
// difference from Python: Go won't compile if you declare a variable and
// never use it. This keeps code clean but surprises beginners.
package main

import "fmt"

func main() {
	// === Explicit type declaration with var ===

	// Use var when you want to be specific about the type,
	// or when you need to declare without assigning a value.
	var name string = "Alice"
	var age int = 25
	var price float64 = 19.99
	var isActive bool = true

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Price:", price)
	fmt.Println("Active:", isActive)

	// === Short declaration with := ===

	// := declares AND assigns with type inferred from the value.
	// This is the most common way to create variables in Go.
	// Compare to Python: city = "Philadelphia" (also inferred, but dynamic).
	// Go's inference is still static — once city is a string, it stays a string.
	city := "Philadelphia"
	score := 95.5
	enrolled := true

	fmt.Println("\nCity:", city)
	fmt.Println("Score:", score)
	fmt.Println("Enrolled:", enrolled)

	// === Multiple variables at once ===

	// Declare and assign two variables in one line.
	// You'll see this with function returns: result, err := someFunction()
	x, y := 10, 20
	fmt.Println("\nx:", x, "y:", y)

	// === Constants ===

	// Constants can't be changed after declaration.
	// Use const for values that should never change (config, limits, labels).
	const maxRetries = 3
	const appName = "Health Monitor"
	fmt.Println("\nApp:", appName, "Max retries:", maxRetries)

	// === The unused variable rule ===

	// Uncomment this line and try to compile — Go will refuse.
	// This is intentional: unused variables are dead code, and Go
	// eliminates dead code at the compiler level.
	// unused := "this will cause a compile error"
}
