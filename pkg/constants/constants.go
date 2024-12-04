package constants

import (
	"fmt"
)

// Constants demonstrates the use of constant declarations in Go.
//
// This function performs the following:
// - Declares a constant for the brand name ("Boat") and prints it.
// - Demonstrates that constants cannot be reassigned (commented-out line shows an error if uncommented).
// - Declares two numerical constants (`num1` and `num2`) that perform division and prints the results.
//
// Example:
//
//	Constants() will print:
//	  Boat
//	  43.47826086956522
//	  50
func Constants() {
	const Brand = "Boat"
	// Error
	// Brand ="Google"
	fmt.Println(Brand)

	const num1 = 1e3 / 23
	fmt.Println(num1)

	const num2 = 1e3 / 20
	fmt.Println(num2)
}
