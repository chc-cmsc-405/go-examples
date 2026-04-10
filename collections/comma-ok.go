// comma-ok.go — Safely checking if a map key exists before using its value.
//
// When you read a missing key from a Go map, you get the zero value (0 for int,
// "" for string, false for bool) — silently, with no error. This is different
// from Python (KeyError) and Java (returns null). The comma-ok idiom is how
// you distinguish "key not found" from "key exists with a zero value."
package main

import "fmt"

func main() {
	scores := map[string]int{
		"Alice": 95,
		"Bob":   87,
	}

	// === The problem: zero values are ambiguous ===

	// Charlie isn't in the map, but Go returns 0 — not an error.
	// Is Charlie's score really 0, or does Charlie not exist?
	fmt.Println("Charlie's score:", scores["Charlie"]) // 0 — but is that real?

	// === The fix: comma-ok ===

	// Two return values: the value, and a boolean (true = key exists).
	score, exists := scores["Charlie"]
	if exists {
		fmt.Println("Charlie's score:", score)
	} else {
		fmt.Println("Charlie not found")
	}

	// === Common shorthand ===

	// Declare and check in one line — the variables are scoped to the if block.
	// This is idiomatic Go and you'll see it everywhere.
	if score, ok := scores["Alice"]; ok {
		fmt.Printf("Alice scored %d\n", score)
	}

	// === Why this matters: real zero values ===

	// A temperature of 0.0 is a legitimate reading. Without comma-ok,
	// you can't tell "sensor reported 0 degrees" from "sensor not found."
	temperatures := map[string]float64{
		"server-1": 72.5,
		"server-2": 0.0, // this is a real reading, not a missing key
	}

	// server-2 exists with value 0.0 — comma-ok confirms it's real
	temp, ok := temperatures["server-2"]
	fmt.Printf("\nserver-2: temp=%.1f, exists=%v\n", temp, ok)

	// server-3 doesn't exist — comma-ok catches it
	temp, ok = temperatures["server-3"]
	fmt.Printf("server-3: temp=%.1f, exists=%v\n", temp, ok)
}
