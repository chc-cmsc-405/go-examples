// count.go — Count occurrences of a substring.
//
// strings.Count returns the number of non-overlapping occurrences of
// a substring. Useful when you want to know how often something appears
// without extracting it. Compared to writing a loop with strings.Index,
// strings.Count is one line and reads as intent rather than mechanism.
package main

import (
	"fmt"
	"strings"
)

func main() {
	// === Count a single character ===

	word := "banana"
	aCount := strings.Count(word, "a")
	fmt.Printf("Word: %q\n", word)
	fmt.Printf("Count of 'a': %d\n\n", aCount)

	// === Count a multi-character substring ===

	sentence := "the the quick the brown the the fox"
	theCount := strings.Count(sentence, "the")
	fmt.Printf("Sentence: %q\n", sentence)
	fmt.Printf("Count of 'the': %d\n\n", theCount)

	// === Case sensitivity ===

	// strings.Count is case-sensitive. "The" and "the" count separately.
	mixed := "The the THE"
	fmt.Printf("Mixed: %q\n", mixed)
	fmt.Printf("Count of 'the' (lowercase): %d\n", strings.Count(mixed, "the"))
	fmt.Printf("Count of 'The' (capital T): %d\n", strings.Count(mixed, "The"))
	fmt.Printf("Count of 'THE' (uppercase): %d\n", strings.Count(mixed, "THE"))
}
