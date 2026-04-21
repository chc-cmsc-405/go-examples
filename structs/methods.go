// methods.go — Attaching behavior to a struct with methods.
//
// A method is a function with a receiver — the type it belongs to. Methods
// live OUTSIDE the struct definition (unlike Java/Python, where they live
// inside the class body). The compiler stitches them together.
//
// The receiver `(s Service)` is like `self` in Python or `this` in Java —
// it's how the method refers to the struct it was called on.
package main

import "fmt"

type Service struct {
	Name string
	URL  string
	Tags []string
}

// === Method definition ===

// Methods are defined outside the struct — this looks strange to Java/Python
// developers but is standard Go. The receiver `(s Service)` binds this
// function to the Service type.
func (s Service) IsTagged(tag string) bool {
	for _, t := range s.Tags {
		if t == tag {
			return true
		}
	}
	return false
}

// Multiple methods on the same type — each one is its own top-level function
// with its own receiver. They don't need to be grouped together in the file.
func (s Service) TagCount() int {
	return len(s.Tags)
}

func main() {
	google := Service{
		Name: "Google",
		URL:  "https://google.com",
		Tags: []string{"search", "work"},
	}

	// === Calling methods ===

	// Method call syntax is identical to Java/Python — dot-notation.
	// The receiver `s` inside IsTagged becomes `google` here.
	fmt.Println("Is 'work' a tag?", google.IsTagged("work"))
	fmt.Println("Is 'streaming' a tag?", google.IsTagged("streaming"))
	fmt.Println("Tag count:", google.TagCount())

	// === Methods on other instances ===

	// Each instance's method sees its own data through the receiver.
	netflix := Service{Name: "Netflix", Tags: []string{"streaming", "fun"}}
	fmt.Println("\nNetflix tagged 'streaming'?", netflix.IsTagged("streaming"))
	fmt.Println("Netflix tagged 'work'?", netflix.IsTagged("work"))
}
