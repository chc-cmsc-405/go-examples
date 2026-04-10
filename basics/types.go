// types.go — Go's basic types: int, float64, bool, string.
//
// If you know Python, Go's type system is the biggest adjustment. Python
// is dynamically typed — a variable can be a string one moment and an int
// the next. Go is statically typed — once a variable is declared as string,
// it stays string forever. Go also requires explicit type conversion
// (no automatic int-to-float). This catches bugs at compile time that
// Python would let through until runtime.
package main

import "fmt"

func main() {
	// === Integer types ===

	// int is the default — platform-sized (64-bit on modern machines).
	// Go also has int8, int16, int32, int64 for specific sizes.
	// Compare to Java: int is always 32-bit; Go's int matches the platform.
	var count int = 42
	var small int8 = 127 // -128 to 127 — useful for small ranges
	fmt.Println("Count:", count)
	fmt.Println("Small:", small)

	// === Floating point ===

	// float64 is Go's default float — like Python's float and Java's double.
	// float32 exists but you'll rarely need it.
	var price float64 = 19.99
	var ratio float32 = 0.75
	fmt.Println("Price:", price)
	fmt.Println("Ratio:", ratio)

	// === Boolean ===
	var isReady bool = true
	fmt.Println("Ready:", isReady)

	// === Strings ===

	// Lowercase 'string' — not 'String' like Java.
	// Go strings are immutable (like Python and Java).
	var greeting string = "Hello, Go!"
	fmt.Println("Greeting:", greeting)
	fmt.Println("Length:", len(greeting))

	// === Explicit type conversion ===

	// Go NEVER converts types automatically. This line would fail:
	//   var z float64 = x   // ERROR: cannot use x (int) as float64
	// You must be explicit — this prevents subtle precision bugs.
	var x int = 10
	var y float64 = float64(x) // explicit conversion required
	fmt.Println("\nInt:", x)
	fmt.Println("Float:", y)

	// === Zero values ===

	// Uninitialized variables get a default "zero value" — not nil, not undefined.
	// This is different from Python (NameError) and Java (null for objects).
	// Every type has a zero value: 0 for numbers, "" for strings, false for bools.
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
