// status-analyzer.go — A realistic StatusAnalyzer that makes a real HTTP request.
//
// Notice the Analyze method ignores its `body` parameter. Interfaces constrain
// the method SIGNATURE, not the implementation. StatusAnalyzer needs to measure
// response time, so it fetches the URL itself rather than using a pre-fetched
// body. Completely fine — the interface only requires that the method exists
// with the right signature.
package main

import (
	"fmt"
	"net/http"
	"time"
)

type AnalysisResult struct {
	AnalyzerName string
	Data         map[string]any
}

type Analyzer interface {
	Analyze(url string, body []byte) AnalysisResult
	Name() string
}

type StatusAnalyzer struct{}

// === body parameter is present but unused ===

// The method has to match the interface signature, so `body` stays in the
// parameter list. Go's compiler doesn't complain about unused parameters
// (only unused local variables).
func (s StatusAnalyzer) Analyze(url string, body []byte) AnalysisResult {
	start := time.Now()
	resp, err := http.Get(url)
	elapsed := time.Since(start)

	data := map[string]any{"response_ms": elapsed.Milliseconds()}
	if err != nil {
		data["error"] = err.Error()
	} else {
		resp.Body.Close()
		data["status_code"] = resp.StatusCode
	}
	return AnalysisResult{AnalyzerName: s.Name(), Data: data}
}

func (s StatusAnalyzer) Name() string { return "status" }

func main() {
	s := StatusAnalyzer{}

	// === Call through the concrete type ===
	result := s.Analyze("https://example.com", nil)
	fmt.Printf("Direct call: %+v\n", result)

	// === Call through the interface ===

	// Same method dispatches — the interface is just a different way of
	// holding the same value. The underlying code runs either way.
	var a Analyzer = s
	result = a.Analyze("https://example.com", nil)
	fmt.Printf("Interface call: %+v\n", result)

	// === Error path — bad URL ===

	// Demonstrates the interface handling failure gracefully. The caller
	// doesn't need to know how StatusAnalyzer reports errors — it's all
	// packed into the Data map.
	result = s.Analyze("http://localhost:1", nil)
	fmt.Printf("\nError case: %+v\n", result)
}
