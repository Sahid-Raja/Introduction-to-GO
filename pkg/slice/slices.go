package slice

import (
	"fmt"
	"slices"
)

func Slice() {
	var s []string
	fmt.Println("Uninatalized ", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("Slice of Zero length ", s, s == nil, len(s), cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("After using Append Function ", s)

	copyString := make([]string, len(s))
	copy(copyString, s)
	fmt.Println("Copy String is : ", copyString)

	l := s[2:5]
	fmt.Println("1st Slice ", l)
	l = s[:3]
	fmt.Println("2nd Slice ", l)
	l = s[2:]
	fmt.Println("3rd Slice ", l)

	tempString1 := []string{"g", "h", "i"}
	fmt.Println("New Temporary String Created ", tempString1)
	tempString2 := []string{"g", "h", "i"}

	if slices.Equal(tempString1, tempString2) {
		fmt.Println("Both the Slices are Equal")
	}

	twoDSLice := make([][]int, 3)
	for i := 0; i < 3; i++ {
		twoDSLice[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			twoDSLice[i][j] = i + j
		}
	}

}
