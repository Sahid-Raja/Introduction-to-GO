package recurssion

import "fmt"

func recurssiveBinarySearch(k, left, right int, arr []int) int {
	if left > right {
		return -1
	}
	mid := left + ((right - left) / 2)
	if arr[mid] == k {
		return mid
	} else if arr[mid] > k {
		return recurssiveBinarySearch(k, left, mid-1, arr)
	} else {
		return recurssiveBinarySearch(k, mid+1, right, arr)
	}
}
func Recurssion() {
	var arr []int
	for i := 0; i < 10; i++ {
		arr = append(arr, i+1)
	}

	var printArray func(int)

	printArray = func(i int) {
		if i == len(arr) {
			println()
			return
		}
		fmt.Print(arr[i], " ")
		i++
		printArray(i)
	}

	printArray(0)
	fmt.Printf("Key is Avilable at index %v", recurssiveBinarySearch(5, 0, len(arr)-1, arr))

}
