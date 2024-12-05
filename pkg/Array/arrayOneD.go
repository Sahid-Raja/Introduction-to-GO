package array

import "fmt"

func twoSum(sum int, arr [5]int) (int, int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] == (sum - arr[i]) {
				return arr[i], arr[j]
			}
		}
	}

	return -1, -1
}

func ArrayOneD() {

	arr := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("Given Array is %v\n", arr)
	sum := 8
	fmt.Printf("Sum to Find is %v\n", sum)

	num1, num2 := twoSum(sum, arr)

	if num1 == -1 || num2 == -1 {
		fmt.Println("Given Sum Can't be formed")
	} else {
		fmt.Printf("Given Sum Can be formed by Using %d and %d\n", num1, num2)
	}

}
