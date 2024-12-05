package closures

import "fmt"

func powerOf(power int) func(int) (int, int) {
	return func(num int) (int, int) {
		res := 1
		for i := 0; i < power; i++ {
			res = res * num
		}
		return num, res
	}
}

func Closures() {

	fmt.Println("Square till 5")
	sqaureOfNum := powerOf(2)
	for i := range 5 {
		num, numSquare := sqaureOfNum(i)
		fmt.Println(num, " ", numSquare)
	}

	fmt.Println("Cube till 5")
	cubeOfNum := powerOf(3)
	for i := range 5 {
		num, numCube := cubeOfNum(i)
		fmt.Println(num, " ", numCube)
	}
}
