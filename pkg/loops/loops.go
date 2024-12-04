package loops

import (
	"fmt"

	days "github.com/Sahid-Raja/Introduction-to-GO/pkg/switch"
)

// Loops demonstrates different loop types in Go.
// - Loops from 0 to 7, calling 'Switch' to print the corresponding day of the week.
// - Loops from 0 to 10 to check if numbers are zero, even, or odd, using 'continue' for zero and bitwise operations for even numbers.
// - Demonstrates a faulty range loop that will result in a compile-time error due to incorrect usage.
func Loops() {

	// Loop through the days and print the day names.
	i := 0
	for i <= 7 {
		fmt.Println(days.Days(i)) // Get day name using 'Switch'
		i++
	}

	// Loop from 0 to 10 and print whether numbers are zero, even, or odd.
	for j := 0; j <= 10; j++ {
		if j == 0 {
			fmt.Println("Number is Zero")
			continue
		}
		if j&1 == 0 {
			fmt.Printf("%v %v", j, "is an Even Number\n")
		}
		fmt.Printf("%v %v", j, "is an odd Number\n")
	}

	for j := range 3 {
		fmt.Print(j, " ")
	}
}
