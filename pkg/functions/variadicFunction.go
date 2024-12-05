package functions

import "fmt"

func tSum(nums ...int) int {
	fmt.Println(nums, " ")
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	return totalSum
}

func VariadicFunctions() {
	tSum(1, 2)
	tSum(1, 2, 3)
	tSum(1, 2, 3, 4)

	nums := []int{1, 2, 3, 4, 5, 6}
	tSum(nums...)
}
