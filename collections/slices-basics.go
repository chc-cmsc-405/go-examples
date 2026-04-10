// slices-basics.go — Creating slices, accessing elements, and basic operations.
//
// Slices are Go's dynamic list — like Python's list or Java's ArrayList.
// Unlike Python, Go also has arrays (fixed size), but you'll almost never
// use them directly. Slices are what you reach for. The key difference
// from Python: Go slices are typed ([]string holds only strings).
package main

import "fmt"

func main() {
	// === Arrays (for comparison — you won't use these much) ===

	// Arrays have a fixed size baked into the type: [3]float64.
	// You can't append to an array — it will never grow.
	var temps [3]float64
	temps[0] = 72.5
	temps[1] = 68.3
	temps[2] = 75.1
	fmt.Println("Array:", temps)
	fmt.Println("Length:", len(temps))

	// === Slices (this is what you'll use) ===

	// No size in the brackets — that's what makes it a slice, not an array.
	// Compare to Python: services = ["Google", "GitHub", "Netflix"]
	services := []string{"Google", "GitHub", "Netflix"}
	fmt.Println("\nSlice:", services)
	fmt.Println("Length:", len(services))

	// Access by index — same as Python and Java.
	fmt.Println("First:", services[0])
	fmt.Println("Last:", services[len(services)-1])

	// Slice literal with numbers
	responseTimes := []int{120, 85, 200, 95, 150}
	fmt.Println("\nResponse times:", responseTimes)
}
