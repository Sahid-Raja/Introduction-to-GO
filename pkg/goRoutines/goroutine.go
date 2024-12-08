package goroutines

import (
	"fmt"
	"time"
)

func printNum(typeOfProcesss string) {
	for i := 0; i <= 3; i++ {
		fmt.Println(typeOfProcesss, " ", i)
	}
}

func Goroutines() {
	printNum("Direct")

	go printNum("Using GO Routines")

	go func(s string) {
		time.Sleep(time.Microsecond * 10)
		fmt.Println("Loading ", s)
	}("Hello WOrld")

	time.Sleep(time.Second)
}
