// hello.go — Your first Go program.
//
// Every Go program starts with a package declaration and a main function.
// package main tells Go "this is an executable program" (not a library).
// import "fmt" loads the formatting package — Go's standard I/O library.
// If you import a package and don't use it, Go won't compile. That's
// intentional — no dead code allowed.
package main

import "fmt"

func main() {
	// Println adds a newline automatically — like Python's print().
	fmt.Println("Hello, World!")
	fmt.Println("Welcome to Go!")
}
