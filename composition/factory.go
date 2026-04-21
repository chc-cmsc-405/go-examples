// factory.go — Factory functions hide awkward construction syntax.
//
// When a type embeds another type, building instances with a struct literal
// gets noisy: StatusAnalyzer{BaseAnalyzer: BaseAnalyzer{AnalyzerName: "status"}}.
// A factory function (Go's convention: NewTypeName) wraps that in a clean
// single call. Same pattern as constructors in Java/Python — just a regular
// function returning a value.
package main

import "fmt"

type BaseAnalyzer struct {
	AnalyzerName string
}

func (b BaseAnalyzer) Name() string { return b.AnalyzerName }

type StatusAnalyzer struct {
	BaseAnalyzer
}

type LinkAnalyzer struct {
	BaseAnalyzer
}

// === Factory functions ===

// Naming convention: New<TypeName>. Returns the struct by value (or *Type
// if the type must be addressable or is large). Keeps the noisy
// nested-literal out of callers' code.
func NewStatusAnalyzer() StatusAnalyzer {
	return StatusAnalyzer{BaseAnalyzer: BaseAnalyzer{AnalyzerName: "status"}}
}

func NewLinkAnalyzer() LinkAnalyzer {
	return LinkAnalyzer{BaseAnalyzer: BaseAnalyzer{AnalyzerName: "links"}}
}

func main() {
	// === Direct struct literal — verbose, error-prone ===

	// You have to spell out the embedded type name both as the outer field
	// and as the inner type. Easy to misspell, easy to get wrong if the
	// AnalyzerName default ever changes.
	verbose := StatusAnalyzer{BaseAnalyzer: BaseAnalyzer{AnalyzerName: "status"}}

	// === Factory function — clean, centralized ===

	// Callers don't need to know the internal structure. If you ever add
	// fields or change defaults, only the factory changes.
	clean := NewStatusAnalyzer()
	other := NewLinkAnalyzer()

	fmt.Println("Verbose literal: ", verbose.Name())
	fmt.Println("Factory call:    ", clean.Name())
	fmt.Println("Second factory:  ", other.Name())

	// === When to add a factory ===

	// Write a factory when construction needs any of:
	//  - default values for fields the caller shouldn't have to supply
	//  - validation (return an error if the input is invalid)
	//  - dependency setup (open a file, connect to a service)
	//  - non-trivial field initialization (embedded types, computed fields)
	//
	// For a simple two-field struct with no logic, the literal is fine —
	// don't write a factory just because you can.
}
