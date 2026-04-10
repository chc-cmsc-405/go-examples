// slice-vs-array.go — The difference between arrays (fixed) and slices (dynamic).
//
// One character distinguishes them: [3]string is an array, []string is a slice.
// Python only has list (dynamic). Java has both String[] and ArrayList<String>.
// Go has both too, but slices are what you'll use 99% of the time. This file
// shows when you'd use each and how make() and slicing syntax work.
package main

import "fmt"

func main() {
	// === Array — fixed size ===

	// The size is part of the type: [3]string is a different type from [4]string.
	// You can't append to an array — it's fixed at creation.
	arr := [3]string{"a", "b", "c"}
	// arr = append(arr, "d")  // Won't compile — arrays don't support append

	// === Slice — dynamic ===

	// No size in the brackets. Slices can grow with append().
	slc := []string{"a", "b", "c"}
	slc = append(slc, "d") // Works — slices grow as needed

	fmt.Println("Array:", arr, "— length:", len(arr))
	fmt.Println("Slice:", slc, "— length:", len(slc))

	// === make() — pre-allocate a slice with capacity ===

	// make(type, length, capacity) creates a slice with room to grow.
	// length = how many items it starts with (0 here — empty).
	// capacity = how many items it can hold before needing to reallocate (10).
	// This is a performance optimization — the code works the same without it.
	results := make([]string, 0, 10)
	results = append(results, "UP")
	results = append(results, "DOWN")
	fmt.Println("\nResults:", results)
	fmt.Printf("Length: %d, Capacity: %d\n", len(results), cap(results))

	// === Slicing syntax — extracting a portion ===

	// Works exactly like Python: [start:end] where start is inclusive, end is exclusive.
	// [:3] means "from the beginning through index 2."
	// [3:] means "from index 3 to the end."
	names := []string{"Alice", "Bob", "Charlie", "Diana", "Eve"}
	first3 := names[:3]
	last2 := names[3:]
	middle := names[1:4]
	fmt.Println("\nAll:", names)
	fmt.Println("First 3:", first3)
	fmt.Println("Last 2:", last2)
	fmt.Println("Middle:", middle)
}
