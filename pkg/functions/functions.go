package functions

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func multiplication(a, b, c int) int {
	return a * b * c
}
func Functions() {
	sumOfTwoNumbers := sum(2, 4)
	multiplicationOfThreeNumbers := multiplication(5, 6, 7)

	fmt.Printf("Sum is %v and Multiplication result is %v ", sumOfTwoNumbers, multiplicationOfThreeNumbers)
}
