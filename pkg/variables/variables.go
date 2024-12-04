package variables

import (
	"fmt"
	"time"

	days "github.com/Sahid-Raja/Introduction-to-GO/pkg/switch"
	// Importing the 'days' package where the 'Switch' function is defined to get the name of the day.
)

// Variables demonstrates variable declarations, printing their values,
// and formatting a message that includes the current day of the week.
//
// It performs the following:
// - Declares variables of different types (integer, floating-point, boolean, and string).
// - Uses the 'Switch' function from the 'days' package to get the day of the week based on the current date.
// - Prints the values of the variables and a formatted greeting message that includes the name and the current day.
//
// Example:
//
//	Variables() will print something like:
//	  4
//	  5
//	  true
//	  Hello! Sahid Raja Ansari and today's day is <current_day>
func Variables() {
	var num int = 4
	var floatNum = 5.0
	isEmpty := true
	println(num)
	println(floatNum)
	println(isEmpty)

	var greet string = "Hello!! "
	var name, dayInt = "Sahid Raja Ansari", time.Now().Day()
	s := fmt.Sprintf(greet + name + " and today's day is " + days.Days(dayInt))
	println(s)
}
