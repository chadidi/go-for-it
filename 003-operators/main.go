// Package main defines the entry point for the Go program.
package main

// Importing the fmt package for formatted I/O operations.
import "fmt"

// main is the entry point of the Go program.
func main() {

	// Arithmetic operators
	a := 10
	b := 5
	fmt.Println("a + b = ", a+b)
	fmt.Println("a - b = ", a-b)
	fmt.Println("a * b = ", a*b)
	fmt.Println("a / b = ", a/b)
	fmt.Println("a mod b = ", a%b)

	// Relational operators
	fmt.Println("a == b:", a == b)
	fmt.Println("a != b:", a != b)
	fmt.Println("a > b:", a > b)
	fmt.Println("a < b:", a < b)
	fmt.Println("a >= b:", a >= b)
	fmt.Println("a <= b:", a <= b)

	// Logical operators
	c := true
	d := false
	fmt.Println("Logical AND:", c && d)
	fmt.Println("Logical OR:", c || d)
	fmt.Println("Logical NOT:", !c)

	// Not covered yet
	// // Bitwise operators
	// // Assignment operators

}
