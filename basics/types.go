package main

import "fmt"

func main() {
	// Integer types
	var count int = 42
	var small int8 = 127 // -128 to 127
	fmt.Println("Count:", count)
	fmt.Println("Small:", small)

	// Floating point
	var price float64 = 19.99
	var ratio float32 = 0.75
	fmt.Println("Price:", price)
	fmt.Println("Ratio:", ratio)

	// Boolean
	var isReady bool = true
	fmt.Println("Ready:", isReady)

	// Strings
	var greeting string = "Hello, Go!"
	fmt.Println("Greeting:", greeting)
	fmt.Println("Length:", len(greeting))

	// Type conversion — Go requires explicit conversion
	var x int = 10
	var y float64 = float64(x) // must convert explicitly
	fmt.Println("\nInt:", x)
	fmt.Println("Float:", y)

	// This would NOT compile — Go doesn't do implicit conversion:
	// var z float64 = x   // ERROR: cannot use x (int) as float64

	// Zero values — uninitialized variables get defaults
	var emptyString string
	var zeroInt int
	var zeroBool bool
	var zeroFloat float64
	fmt.Println("\nZero values:")
	fmt.Printf("  string: %q\n", emptyString) // ""
	fmt.Printf("  int: %d\n", zeroInt)         // 0
	fmt.Printf("  bool: %t\n", zeroBool)       // false
	fmt.Printf("  float: %f\n", zeroFloat)     // 0.000000
}
