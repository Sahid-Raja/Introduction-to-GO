package array

import "fmt"

func pascalTriangle() {

	var arr [5][5]int

	for i := 0; i < 5; i++ {
		arr[i][0] = 1
		arr[i][i] = 1
		for j := 1; j < i; j++ {
			arr[i][j] = arr[i-1][j-1] + arr[i-1][j]
			// fmt.Printf("Index are %d and %d and Value is %d after adding %d, %d\n", i, j, arr[i][j], arr[i-1][j-1], arr[i-1][j])
		}
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%d %v", arr[i][j], " ")
		}
		fmt.Println()
	}
}
func ArrayTwoD() {
	// var twoD [2][3]int
	// for i := 0; i < 2; i++ {
	//     for j := 0; j < 3; j++ {
	//         twoD[i][j] = i + j
	//     }
	// }
	// fmt.Println("2d: ", twoD)

	// twoD = [2][3]int{
	//     {1, 2, 3},
	//     {1, 2, 3},
	// }
	// fmt.Println("2d: ", twoD)
	pascalTriangle()
}
