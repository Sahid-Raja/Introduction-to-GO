package strings

import (
	"fmt"
	"unicode/utf8"
)

func Strings() {
	var s = "IAmआम"
	fmt.Println("Len of s is: ", len(s))
	for i := range len(s) {
		fmt.Printf("%d ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune Count: ", utf8.RuneCountInString(s))

	for index, runeElement := range s {
		fmt.Printf("%v starts at %d\n", string(runeElement), index)
	}

	//TODO:
	// fmt.Println([]byte(s))
}
