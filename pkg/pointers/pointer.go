package pointers

import "fmt"

func callByValue(num int) {
	num = 1
}
func callByAddress(num *int) {
	*num = 1
}

func checkingForSlice(sliceArray []int) {
	sliceArray[len(sliceArray)-1] = sliceArray[len(sliceArray)-1] + 1
}
func Pointers() {
	num := 2
	fmt.Println("Initial Value is ", num)

	callByValue(num)
	fmt.Println("After Calling Call by value Value is ", num)
	callByAddress(&num)
	fmt.Println("After Calling Call by address Value is ", num)

	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println("Before Calling Function ", arr)
	checkingForSlice(arr[:])
	fmt.Println("After Calling Function ", arr)

}
