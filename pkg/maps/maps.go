package maps

import (
	"fmt"

	swtch "github.com/Sahid-Raja/Introduction-to-GO/pkg/switch"
)

func Maps() {
	uMap := make(map[int]string)

	for i := 1; i <= 7; i++ {
		uMap[i] = swtch.Days(i)
	}
	fmt.Println("Map after Intialization ", uMap)

	uMap[0] = "No Day"
	fmt.Println("Map after Inserting new Element ", uMap)

	delete(uMap, 3)
	delete(uMap, 4)
	fmt.Println("Map after deleting 3 and 4 keys ", uMap)

	// clear(uMap)
	// fmt.Println("After clearing the Map ", uMap)

	value, isAvilable := uMap[1]
	fmt.Printf("%v -> %v \n", value, isAvilable)
	_, isAvilable = uMap[3]
	fmt.Printf("%v \n", isAvilable)

	newMap := map[string]string{"fruit": "mango", "animal": "girrafe", "food": "indian"}
	fmt.Print("New Map is ", newMap)
}
