// pointer-receivers.go — Value receivers vs pointer receivers.
//
// This is the most-tested idea in Phase 3. A value receiver (s Service)
// operates on a COPY of the struct — mutations are thrown away. A pointer
// receiver (s *Service) operates on the ORIGINAL — mutations stick.
//
// Rule of thumb: any method that mutates state needs a pointer receiver.
// Any method that's read-only can use a value receiver. If any method on the
// type needs a pointer receiver, use pointer receivers for all methods on
// that type (consistency avoids subtle bugs).
package main

import "fmt"

type Service struct {
	Name string
	URL  string
}

// === Version A: value receiver — DOES NOT mutate the original ===

// `s` is a copy of the Service. Any changes to s.Name affect only the copy,
// which is discarded when the method returns.
func (s Service) RenameValue(newName string) {
	s.Name = newName // changes the copy, not the original
}

// === Version B: pointer receiver — DOES mutate the original ===

// `s *Service` is a pointer to the original. s.Name modifies the struct
// the caller holds, so the change is visible after the method returns.
// Go auto-dereferences — you write s.Name, not (*s).Name.
func (s *Service) RenamePointer(newName string) {
	s.Name = newName // changes the original
}

func main() {
	google := Service{Name: "Google", URL: "https://google.com"}

	// === Value receiver — the mutation is lost ===

	fmt.Println("Before RenameValue:", google.Name)
	google.RenameValue("Alphabet")
	fmt.Println("After RenameValue: ", google.Name) // Still "Google"

	// === Pointer receiver — the mutation sticks ===

	fmt.Println("\nBefore RenamePointer:", google.Name)
	google.RenamePointer("Alphabet")
	fmt.Println("After RenamePointer: ", google.Name) // Now "Alphabet"

	// === Why Go requires you to be explicit ===

	// Java hides this — every method on a class can mutate any field.
	// Go makes you declare your intent in the method signature. You can tell
	// at a glance whether a method will mutate its receiver: look for the *.
	//
	// This is more verbose but removes a whole category of "oh, that method
	// silently changed my object" bugs.
}
