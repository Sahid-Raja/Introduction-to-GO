package values

import "fmt"

// Values demonstrates the usage of various data types and logical operations in Go.
//
// This function performs the following actions:
// - Defines variables of different types (string, int8, float64).
// - Prints the concatenated first name and last name.
// - Prints an 8-bit integer.
// - Performs and prints a floating-point division.
// - Demonstrates logical AND, OR, and NOT operations.
//
// Example:
//   Values() prints:
//     Sahid Raja Ansari
//     127
//     3.5714285714285716
//     false
//     true
//     true
func Values() {
	var firstName string = "Sahid Raja"
	var lastName string = "Ansari"
	var eightBitNum int8 = 127
	floatNum := 25.0 / 7.0

	// Output the concatenated full name.
	fmt.Println(firstName + " " + lastName)

	// Output the 8-bit integer value.
	fmt.Println(eightBitNum)

	// Output the result of the floating-point division.
	fmt.Println(floatNum)

	// Output the result of a logical AND operation (false).
	fmt.Println(true && false)

	// Output the result of a logical OR operation (true).
	fmt.Println(true || false)

	// Output the result of a logical NOT operation (true).
	fmt.Println(!false)
}
