package swtch

import (
	"fmt"
	"reflect"
)

func SwitchFunction() {
	var x = 112.2
	switch {
	case reflect.TypeOf(x).String() == "int":
		fmt.Printf("Type of x is integer")
	case reflect.TypeOf(x).String() == "string":
		fmt.Printf("Type of x is String")
	case reflect.TypeOf(x).String() == "bool":
		fmt.Printf("Type of x is Boolean")
	case reflect.TypeOf(x).String() == "float64":
		fmt.Printf("Type of x is float64")
	default:
		fmt.Println("Type not Known")
	}
}
