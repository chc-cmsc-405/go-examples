// embed.go — Composition through struct embedding.
//
// Embedding is Go's replacement for inheritance. Put a struct inside another
// struct WITHOUT a field name, and the inner struct's fields and methods are
// "promoted" to the outer struct — they appear to belong to the outer type.
//
// Critical distinction: embedding is HAS-A, not IS-A. A StatusAnalyzer HAS
// a BaseAnalyzer; it IS NOT a BaseAnalyzer. Functions that take BaseAnalyzer
// will not accept StatusAnalyzer. Embedding gives you code reuse; interfaces
// give you polymorphism. Go keeps these concerns separate on purpose.
package main

import "fmt"

// === The shared base ===

type BaseAnalyzer struct {
	AnalyzerName string
}

func (b BaseAnalyzer) Name() string {
	return b.AnalyzerName
}

// === Two types that embed BaseAnalyzer ===

// StatusAnalyzer has BaseAnalyzer as an embedded field — no field name.
// Its fields (AnalyzerName) and methods (Name) are promoted.
type StatusAnalyzer struct {
	BaseAnalyzer
	// other StatusAnalyzer-specific fields could go here
}

type LinkAnalyzer struct {
	BaseAnalyzer
}

func main() {
	// === Construction — nested literal ===

	// The outer struct's literal includes the embedded struct by its TYPE
	// NAME as the field name. Awkward at first, but consistent.
	s := StatusAnalyzer{
		BaseAnalyzer: BaseAnalyzer{AnalyzerName: "status"},
	}
	l := LinkAnalyzer{
		BaseAnalyzer: BaseAnalyzer{AnalyzerName: "links"},
	}

	// === Method promotion ===

	// Neither StatusAnalyzer nor LinkAnalyzer defines a Name() method.
	// Both calls work because Name() is promoted from BaseAnalyzer.
	fmt.Println("StatusAnalyzer.Name():", s.Name())
	fmt.Println("LinkAnalyzer.Name(): ", l.Name())

	// === Field promotion works too ===

	// s.AnalyzerName is really s.BaseAnalyzer.AnalyzerName. Go lets you
	// skip the inner field name when there's no ambiguity.
	fmt.Println("\ns.AnalyzerName (promoted):", s.AnalyzerName)

	// === HAS-A, not IS-A ===

	// Uncomment the line below to see the compile error. A function that
	// takes BaseAnalyzer won't accept a StatusAnalyzer — embedding does
	// NOT create a subtype relationship.
	//
	// printName(s)              // compile error: cannot use s as BaseAnalyzer
	printName(s.BaseAnalyzer) // works — pass the embedded field explicitly

	// If you want polymorphism (treat different types the same way),
	// use an interface — that's a separate tool from embedding.
}

func printName(b BaseAnalyzer) {
	fmt.Println("printName got:", b.Name())
}
