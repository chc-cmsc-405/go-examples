// maps-basics.go — Creating, reading, updating, and deleting map entries.
//
// Maps are Go's key-value data structure — the equivalent of Python's dict
// or Java's HashMap. You'll use them constantly: storing services by name,
// counting tags, tracking status. Unlike Python, Go maps require you to
// declare the key and value types up front (map[string]int).
package main

import "fmt"

func main() {
	// === Creating a map ===

	// make() allocates an empty map. You must use make() or a literal —
	// a nil map will panic if you try to write to it.
	scores := make(map[string]int)

	// Adding entries — bracket syntax, same as Python dicts.
	// No .put() like Java — just assign directly.
	scores["Alice"] = 95
	scores["Bob"] = 87
	scores["Charlie"] = 92

	fmt.Println("All scores:", scores)
	fmt.Println("Alice's score:", scores["Alice"])

	// === Updating and deleting ===

	// Update uses the same syntax as add — if the key exists, it overwrites.
	scores["Bob"] = 90
	fmt.Println("Bob's updated score:", scores["Bob"])

	// delete() is a built-in function — no return value, no error if key is missing.
	// Compare to Python's del d["key"] (raises KeyError if missing) or
	// Java's map.remove("key") (returns the old value).
	delete(scores, "Charlie")
	fmt.Println("After deleting Charlie:", scores)

	// === Map literal ===

	// When you know the data at declaration time, use a literal instead of make().
	// Notice the trailing comma after the last entry — Go requires it.
	status := map[string]string{
		"google.com": "UP",
		"github.com": "UP",
		"fake.xyz":   "DOWN",
	}
	fmt.Println("\nService status:", status)

	// len() works on maps just like on slices and strings.
	fmt.Printf("Monitoring %d services\n", len(status))
}
