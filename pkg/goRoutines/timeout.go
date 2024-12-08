package goroutines

import (
	"fmt"
	"time"
)

func TimeOut() {
	done := make(chan bool)

	go func() {
		time.Sleep(1 * time.Second)
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("Program is completed")
	case <-time.After(2 * time.Second):
		fmt.Println("Program timeout")
	}

	// fmt.Println(time.Now().String())

	// timer := time.After(3 * time.Second)

	// fmt.Println(<-timer)
}
