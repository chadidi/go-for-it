// Package main defines the entry point for the Go program.
package main

// Importing the fmt package for formatted I/O operations.
import "fmt"

// main is the entry point of the Go program.
func main() {

	x := 9

	fmt.Println("Variable:", x)
	fmt.Println("Pointer:", &x)

	// This won't change the value of X
	wontChangeValue(x)
	fmt.Println("Value:", x)

	// Pass the pointer to change the value of X
	willChangeValue(&x)
	fmt.Println("Value:", x)
}

func wontChangeValue(x int) {
	x = 11
}

func willChangeValue(x *int) {
	*x = 11
}
