// analyzer.go — Defining an interface and satisfying it implicitly.
//
// Go's interfaces are the biggest departure from Java's OOP: there is no
// `implements` keyword. A type satisfies an interface just by having the
// right methods. The compiler checks the match at the USE site (assignment,
// function call), not at the type definition.
//
// This is called "structural typing" — the structure of the type (its
// method set) determines which interfaces it satisfies.
package main

import "fmt"

// === The result type every analyzer returns ===

// All analyzers return the same shape: their own name + a flexible data map.
// The map uses `any` (alias for interface{}) because different analyzers
// return different shapes of data — status codes, link lists, word counts.
type AnalysisResult struct {
	AnalyzerName string
	Data         map[string]any
}

// === The interface ===

// An interface is a list of method signatures. Any type with these methods
// satisfies the interface. No declaration needed on the implementer.
type Analyzer interface {
	Analyze(url string, body []byte) AnalysisResult
	Name() string
}

// === First implementer — StatusAnalyzer ===

// No `implements Analyzer` anywhere. StatusAnalyzer just has the right methods.
type StatusAnalyzer struct{}

func (s StatusAnalyzer) Analyze(url string, body []byte) AnalysisResult {
	return AnalysisResult{
		AnalyzerName: "status",
		Data:         map[string]any{"body_bytes": len(body)},
	}
}

func (s StatusAnalyzer) Name() string { return "status" }

// === Second implementer — PingAnalyzer ===

// Completely different code. Same interface satisfaction, just by having
// an Analyze and a Name method with matching signatures.
type PingAnalyzer struct{}

func (p PingAnalyzer) Analyze(url string, body []byte) AnalysisResult {
	return AnalysisResult{
		AnalyzerName: "ping",
		Data:         map[string]any{"reachable": true},
	}
}

func (p PingAnalyzer) Name() string { return "ping" }

func main() {
	// === Interface satisfaction is checked at the USE site ===

	// Assigning a StatusAnalyzer into an Analyzer variable is where the
	// compiler checks the method set. If a method is missing or has the
	// wrong signature, this line fails to build.
	var a Analyzer = StatusAnalyzer{}
	fmt.Println("Assigned StatusAnalyzer to Analyzer:", a.Name())

	// === The blank-identifier trick for a compile-time check ===

	// `var _ Type = Value` is a common Go idiom: it compiles if and only if
	// Value satisfies Type. No runtime cost, no unused-variable warning.
	var _ Analyzer = PingAnalyzer{}
	fmt.Println("PingAnalyzer also satisfies Analyzer (compile-time verified)")

	// === Calling through the interface ===

	result := a.Analyze("https://example.com", []byte("hello"))
	fmt.Printf("\n%+v\n", result)
}
