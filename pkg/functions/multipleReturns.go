package functions

import "fmt"

func helper() (int, string) {
	return 1, "Sahid Raja Ansari"
}

func MultipleFunctions() {
	value, name := helper()
	fmt.Println(value, name)
}
