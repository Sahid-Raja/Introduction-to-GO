package errorhandling

import (
	"errors"
	"fmt"
)

var total_storage = 1024
var ErrlowStorageError = fmt.Errorf("storage is full")

func checkTemp(currentTemp int) (int, error) {
	if currentTemp >= 89 {
		return -1, errors.New("device temp is too high")
	} else if currentTemp <= 24 {
		return -1, errors.New("device is currentlty offline")
	} else {
		return currentTemp, nil
	}
}

func checkStorage(size_of_program int) (avilable_storage int, err error) {
	avilable_storage = total_storage - size_of_program
	if avilable_storage < 0 {
		return -1, ErrlowStorageError
	} else {
		return avilable_storage, nil
	}
}

func ErrorHandling() {
	// _, tempError := checkTemp(100)
	for index, element := range []int{10, 20, 30, 40, 50, 30, 52, 70, 100} {
		if currentTemp, err := checkTemp(element); err != nil {
			fmt.Println("failed to connect due to ", err)
		} else {
			fmt.Println("Successfully connected!! current temperature of Device is ", currentTemp, " 'c")
			fmt.Println("Loading Program..........")
			if cs, err := checkStorage(256*index + element); err != nil {
				if errors.Is(err, ErrlowStorageError) {
					fmt.Println("Storage not Aviable please upgrade your ram")
				}
			} else {
				fmt.Printf("Device has %d storage left \n", cs)
			}
		}
	}

}
