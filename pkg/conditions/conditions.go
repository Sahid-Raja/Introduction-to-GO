package conditions

import "fmt"

// The function `Conditions` takes two inputs:
// - `age` (an integer representing the person's age)
// - `gender` (a string representing the person's gender, which can be "Male" or "Female")
//
// Example:
// For a person with age 20 and gender "Female", the function will check:
// - If the person is eligible to vote (age >= 18).
// - If the person is eligible for marriage (age >= 18 for females).
// It will also check if the person is eligible for a learner's license (age >= 16).
func Conditions(age int, gender string) {
	if age >= 18 {
		fmt.Println("You are Eligible to vote")
		if age >= 21 && gender == "Male" {
			fmt.Println("You are eligible to Marriage as well")
		} else if gender == "Female" && age >= 18 {
			fmt.Println("You are eligible to Marriage as well")
		}
	} else {
		fmt.Println("You are not Eligible to vote")
	}

	if tempAge := 16; tempAge >= 16 {
		fmt.Println("You are Eligible for Learner's Liscense")
	} else {
		fmt.Println("You are not Eligible to for Learner's Liscense")
	}
}
