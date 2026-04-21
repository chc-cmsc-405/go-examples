// run-all.go — Driving multiple analyzers behind one interface.
//
// This file is the skeleton your health monitor's /analyze and /analyze-all
// endpoints build on: a []Analyzer slice, a loop, one call per analyzer.
// Adding a fourth analyzer is a two-line change (write the type, append it).
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

// === Three stub analyzers — replace bodies with real logic in the Practice Lab ===

type StatusAnalyzer struct{}

func (s StatusAnalyzer) Analyze(url string, body []byte) AnalysisResult {
	return AnalysisResult{
		AnalyzerName: s.Name(),
		Data:         map[string]any{"stub": "status logic goes here"},
	}
}
func (s StatusAnalyzer) Name() string { return "status" }

type LinkAnalyzer struct{}

func (l LinkAnalyzer) Analyze(url string, body []byte) AnalysisResult {
	return AnalysisResult{
		AnalyzerName: l.Name(),
		Data:         map[string]any{"stub": "link extraction goes here"},
	}
}
func (l LinkAnalyzer) Name() string { return "links" }

type ContentAnalyzer struct{}

func (c ContentAnalyzer) Analyze(url string, body []byte) AnalysisResult {
	return AnalysisResult{
		AnalyzerName: c.Name(),
		Data:         map[string]any{"stub": "title + word count goes here"},
	}
}
func (c ContentAnalyzer) Name() string { return "content" }

// === The driver — one loop runs them all ===

func runAll(url string, body []byte) []AnalysisResult {
	analyzers := []Analyzer{
		StatusAnalyzer{},
		LinkAnalyzer{},
		ContentAnalyzer{},
	}

	results := make([]AnalysisResult, 0, len(analyzers))
	for _, a := range analyzers {
		results = append(results, a.Analyze(url, body))
	}
	return results
}

func main() {
	body := []byte(`<html><title>Hi</title><a href="https://go.dev">Go</a></html>`)
	results := runAll("https://example.com", body)

	for _, r := range results {
		fmt.Printf("%s → %+v\n", r.AnalyzerName, r.Data)
	}
}
