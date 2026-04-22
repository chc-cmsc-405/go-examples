// fields.go — Split on any whitespace, drop empties.
//
// strings.Fields splits on runs of any whitespace character (spaces,
// tabs, newlines) and drops empty strings from the result. This is
// different from strings.Split(s, " "), which only splits on single
// spaces and leaves empty strings around double spaces, tabs, or
// newlines. Use Fields when you want tokens, not pieces.
package main

import (
	"fmt"
	"strings"
)

func main() {
	// A sentence with mixed whitespace: double spaces, a tab, a newline.
	sentence := "The  quick\tbrown\nfox jumps"

	// === strings.Fields drops empties and handles all whitespace ===

	words := strings.Fields(sentence)
	fmt.Printf("Sentence: %q\n", sentence)
	fmt.Printf("Fields:   %q\n", words)
	fmt.Printf("Word count: %d\n\n", len(words))

	// === Compare to strings.Split on a single space ===

	// Split only recognizes the exact delimiter (" "). Double spaces
	// produce empty strings. Tabs and newlines are treated as regular
	// characters inside a token.
	split := strings.Split(sentence, " ")
	fmt.Printf("Split on ' ': %q\n", split)
	fmt.Printf("Split count (with empties and tab/newline artifacts): %d\n", len(split))
}
