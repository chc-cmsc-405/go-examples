// index.go — Find the position of a substring.
//
// strings.Index returns the byte index of the first occurrence of a
// substring, or -1 if the substring is not present. Combine it with
// slice syntax (s[:i], s[i+len(sub):]) to extract what comes before
// or after the match.
package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "The quick brown fox jumps over the lazy dog"

	// === Find a substring ===

	i := strings.Index(sentence, "brown")
	fmt.Printf("Sentence: %q\n", sentence)
	fmt.Printf("Index of 'brown': %d\n\n", i)

	// === Extract what comes before the match ===

	before := sentence[:i]
	fmt.Printf("Before 'brown': %q\n", before)

	// === Extract what comes after the match ===

	// Skip past the matched substring itself with i + len("brown").
	after := sentence[i+len("brown"):]
	fmt.Printf("After 'brown':  %q\n\n", after)

	// === Extract what is between two markers ===

	// A common pattern: find the start marker, find the end marker,
	// take the slice between them. The value of `start` skips past
	// the opening marker so the slice is just the content.
	text := "log: [INFO] server ready"
	open := strings.Index(text, "[")
	close := strings.Index(text, "]")
	inside := text[open+1 : close]
	fmt.Printf("Text: %q\n", text)
	fmt.Printf("Between '[' and ']': %q\n\n", inside)

	// === Missing substring returns -1 ===

	j := strings.Index(sentence, "zebra")
	fmt.Printf("Index of 'zebra': %d\n", j)
	if j == -1 {
		fmt.Println("'zebra' not found — always check for -1 before slicing.")
	}
}
