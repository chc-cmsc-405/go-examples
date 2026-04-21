// basics.go — Defining structs, creating instances, and using factory functions.
//
// Structs are Go's named, typed data containers — the replacement for the
// map[string]interface{} pattern from Phase 2. A struct has fixed fields with
// fixed types, and the compiler checks every access. No more type assertions,
// no more runtime "wrong type" panics.
//
// Compare to: Python dataclasses (@dataclass), Java classes without methods,
// C structs. Go is closer to dataclasses — a data definition, not a behavior
// container. Behavior attaches separately through methods.
package main

import "fmt"

// === Defining a struct ===

// Capitalization controls visibility: Name/URL/Tags are exported (public),
// a lowercase field would be package-private. No public/private keywords.
type Service struct {
	Name string
	URL  string
	Tags []string
}

// === Factory function — Go's constructor convention ===

// Factory functions are named NewTypeName. They encapsulate construction logic
// (defaults, validation) and hide the struct-literal syntax from callers.
func NewService(name, url string, tags []string) Service {
	return Service{Name: name, URL: url, Tags: tags}
}

func main() {
	// === Creating a struct instance ===

	// Named-literal form — the idiomatic default.
	// Field names make the construction self-documenting and protect you if
	// fields are reordered later.
	google := Service{
		Name: "Google",
		URL:  "https://google.com",
		Tags: []string{"search", "work"},
	}

	// Positional form — shorter but fragile. Order must exactly match the
	// struct definition. Don't use this for structs with more than 2 fields.
	// github := Service{"GitHub", "https://github.com", []string{"dev", "work"}}

	// Factory function — how you'll usually build these in real code.
	netflix := NewService("Netflix", "https://netflix.com", []string{"streaming", "fun"})

	// === Field access — no type assertions, no %.(string) ===

	// Compare to Phase 2: service["url"].(string). Direct field access is
	// checked by the compiler — misspell .Nmae and it won't build.
	fmt.Println("Name:", google.Name)
	fmt.Println("URL:", google.URL)
	fmt.Println("Tags:", google.Tags)

	// %+v prints field names alongside values — useful for debugging.
	fmt.Printf("\nFull struct: %+v\n", netflix)

	// === The zero value ===

	// Declare a struct without initializing — every field gets its zero value.
	// No nil-pointer panics, no NullPointerException. Always usable.
	var empty Service
	fmt.Printf("\nZero-value Service: %+v\n", empty)
	fmt.Println("Zero-value Tags is nil but safe to range over:", len(empty.Tags))
}
