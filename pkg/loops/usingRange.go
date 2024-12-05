package loops

import "fmt"

func RangeS() {
	nums := []int{123, 3, 34, 245, 3, 1, 342}
	tSum := 0

	for _, num := range nums {
		tSum += num
	}
	fmt.Println("Total Sum of Array is ", tSum)

	uMap := map[string]string{"Sahid": "Ansari", "John": "Doe", "Alice": "Bob"}

	for key, value := range uMap {
		fmt.Println(key, " ", value)
	}
	println()

	for onlyKey := range uMap {
		fmt.Println(onlyKey)
	}
	println()

	for _, onlyValue := range uMap {
		fmt.Println(onlyValue)
	}
	println()

	for index, character := range "Sahid" {
		fmt.Println(index, " ", string(character))
	}
	println()
}
