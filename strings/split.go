// split.go — Break a string into pieces on a delimiter.
//
// strings.Split returns a []string. Every piece between delimiters
// becomes an element — including empty strings at the start or end
// of the input. The first piece is always whatever came before the
// first delimiter, even if that means everything before the first
// match. This matters when you split text that has a prefix before
// the first interesting piece.
package main

import (
	"fmt"
	"strings"
)

func main() {
	// === Basic split on a comma ===

	// A simple CSV-style line split into fields.
	line := "alice,42,admin"
	fields := strings.Split(line, ",")
	fmt.Printf("Line: %q\n", line)
	fmt.Printf("Pieces: %q\n", fields)
	fmt.Printf("Count: %d\n\n", len(fields))

	// === Ranging over the pieces ===

	// The result of Split is just a []string. Range over it like any slice.
	fmt.Println("Each field:")
	for i, f := range fields {
		fmt.Printf("  [%d] %q\n", i, f)
	}

	// === Split when the input has a prefix before the first delimiter ===

	// If the input has content before the first delimiter, that content
	// becomes pieces[0]. It is NOT a data item — it is the prefix.
	// Code that treats every piece as a data item will misread this.
	log := "START|first|second|third"
	pieces := strings.Split(log, "|")
	fmt.Printf("\nLog: %q\n", log)
	fmt.Printf("Pieces: %q\n", pieces)
	fmt.Println("Note: pieces[0] is \"START\" (the prefix). pieces[1:] are the data items.")
}
