// poly.go — Polymorphism through interfaces, no inheritance required.
//
// Once multiple types satisfy the same interface, you can put them in a
// slice and treat them uniformly. The loop doesn't know or care which
// concrete type it's calling — each call dispatches to the right method
// for the underlying type. That's polymorphism.
//
// Compare to Java: List<Analyzer> does the same thing, but Java requires
// `implements Analyzer` on every class. Go just needs the methods.
package main

import "fmt"

type AnalysisResult struct {
	AnalyzerName string
	Data         map[string]any
}

type Analyzer interface {
	Analyze(url string, body []byte) AnalysisResult
	Name() string
}

type StatusAnalyzer struct{}

func (s StatusAnalyzer) Analyze(url string, body []byte) AnalysisResult {
	return AnalysisResult{AnalyzerName: "status", Data: map[string]any{"code": 200}}
}
func (s StatusAnalyzer) Name() string { return "status" }

type PingAnalyzer struct{}

func (p PingAnalyzer) Analyze(url string, body []byte) AnalysisResult {
	return AnalysisResult{AnalyzerName: "ping", Data: map[string]any{"reachable": true}}
}
func (p PingAnalyzer) Name() string { return "ping" }

type LengthAnalyzer struct{}

func (l LengthAnalyzer) Analyze(url string, body []byte) AnalysisResult {
	return AnalysisResult{AnalyzerName: "length", Data: map[string]any{"bytes": len(body)}}
}
func (l LengthAnalyzer) Name() string { return "length" }

func main() {
	// === Three different types in one slice ===

	// []Analyzer holds any type that satisfies Analyzer. At runtime each slot
	// remembers the concrete type it holds, so method calls dispatch correctly.
	analyzers := []Analyzer{
		StatusAnalyzer{},
		PingAnalyzer{},
		LengthAnalyzer{},
	}

	// === One loop, three different implementations called ===

	// The loop body is uniform. The work each iteration does is not —
	// StatusAnalyzer.Analyze runs for the first element, PingAnalyzer.Analyze
	// for the second, LengthAnalyzer.Analyze for the third.
	url := "https://example.com"
	body := []byte("hello world")

	for _, a := range analyzers {
		result := a.Analyze(url, body)
		fmt.Printf("%s → %+v\n", a.Name(), result.Data)
	}

	// === The design payoff ===

	// Adding a fourth analyzer is a two-line change: write the new type with
	// an Analyze/Name method, append it to the slice. Nothing in this loop
	// or any caller changes. That's the extensibility interfaces give you.
}
