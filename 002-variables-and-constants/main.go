// Package main defines the entry point for the Go program.
package main

// Importing the fmt package for formatted I/O operations.
import "fmt"

// main is the entry point of the Go program.
func main() {

	// Integer data type
	var i int = 42
	fmt.Println("Postive Integer:", i)

	var in int = -42
	fmt.Println("Negative Integer:", in)

	// Unsigned integer data type
	var ui uint = 100
	fmt.Println("Unsigned Integer:", ui)

	// Floating point data type
	var f float64 = 3.14
	fmt.Println("Float:", f)

	// Constant float data type
	const pi float64 = 3.141592653589793238462643383279502884197
	fmt.Println("Pi:", pi) // Notice: only this part (3.141592653589793) was printed not the full pi const

	// String data type
	var s string = "Hello, Go!"
	fmt.Println("String:", s)

	// Boolean data type
	var decolonizePalestine bool = true
	fmt.Println("Boolean:", decolonizePalestine) // https://decolonizepalestine.com/

	// Two variables together
	// Normal
	var (
		xa = 11
		ya = 12
	)

	// Short
	x, y := 11, 12

	fmt.Println("Two variables together - Normal:", xa, ",", ya)
	fmt.Println("Two variables together - Short:", x, ",", y)

	// Not covered yet!
	// // Complex number data type
	// // Array data type
	// // Slice data type
	// // Map data type
	// // Struct data type
	// // Pointer data type

}
