// Package main defines the entry point for the Go program.
package main

// Importing the fmt package for formatted I/O operations.
import "fmt"

// main is the entry point of the Go program.
func main() {
	// More About this Topic Here: https://pkg.go.dev/fmt
	// And Here: https://www.geeksforgeeks.org/fmt-printf-function-in-golang-with-examples/

	// Printing a string with a newline character at the end.
	fmt.Println("Hello, World!")

	// Printing a formatted string using Printf.
	name := "Abdellah"
	age := 27
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	// Printing a floating-point number with two decimal places.
	price := 9.99
	fmt.Printf("The price is $%.2f\n", price)

	// Printing a value in different formats.
	number := 255
	fmt.Printf("Decimal: %d\n", number)
	fmt.Printf("Binary: %b\n", number)
	fmt.Printf("Hexadecimal: %x\n", number)

	// Printing a boolean value.
	isAvailable := true
	fmt.Printf("Available: %t\n", isAvailable)

	// Printing a pointer address.
	ptr := &number
	fmt.Printf("Pointer: %p\n", ptr)

	// Printing a quoted string.
	message := "Bye, Gopher!"
	fmt.Printf("Quoted string: %q\n", message)
}
