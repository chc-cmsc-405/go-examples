// append-range.go — Growing slices with append, iterating with range, filtering.
//
// The biggest gotcha in Go slices: append() returns a new slice. You MUST
// reassign the result — tags = append(tags, "new"). If you forget the
// reassignment, the item is silently lost. Python and Java mutate the list
// in place; Go doesn't. This one difference causes more beginner bugs than
// any other slice behavior.
package main

import "fmt"

func main() {
	// === Append — must reassign ===

	// Start with an empty slice (nil — zero value for slices).
	var tags []string

	// Each append returns a new slice. Assign it back.
	tags = append(tags, "streaming")
	tags = append(tags, "work")
	tags = append(tags, "fun")
	fmt.Println("Tags:", tags)

	// Append multiple items at once with the ... spread operator.
	more := []string{"social", "paid"}
	tags = append(tags, more...)
	fmt.Println("Extended:", tags)

	// === Range — iterating over slices ===

	// range gives you (index, value) — like Python's enumerate().
	fmt.Println("\nAll tags:")
	for i, tag := range tags {
		fmt.Printf("  [%d] %s\n", i, tag)
	}

	// Use _ to ignore the index when you only need the value.
	// Compare to Python: for tag in tags
	fmt.Println("\nJust tags:")
	for _, tag := range tags {
		fmt.Println(" ", tag)
	}

	// === Filtering pattern ===

	// No .filter() method in Go — you write the loop explicitly.
	// This builds a new slice containing only services that match the tag.
	// Compare to Python: [s for s in services if "streaming" in tags[s]]
	services := []string{"Google", "Netflix", "GitHub", "Slack", "Hulu"}
	serviceTags := map[string][]string{
		"Google":  {"search", "work"},
		"Netflix": {"streaming", "fun"},
		"GitHub":  {"dev", "work"},
		"Slack":   {"messaging", "work"},
		"Hulu":    {"streaming", "fun"},
	}

	// Filter: find all services tagged "streaming"
	var streaming []string
	for _, name := range services {
		for _, t := range serviceTags[name] {
			if t == "streaming" {
				streaming = append(streaming, name)
				break // found the tag — stop checking this service's other tags
			}
		}
	}
	fmt.Println("\nAll services:", services)
	fmt.Println("Streaming services:", streaming)
}
