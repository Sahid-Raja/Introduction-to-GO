package goroutines

import "fmt"

func Channel() {
	messages := make(chan string, 3)
	go func() {
		messages <- "Ping me When you Wakeup"
		messages <- "Wake Up"
	}()

	msg := <-messages
	fmt.Println(msg)
	msg = <-messages
	fmt.Println(msg)
	msg = <-messages
	fmt.Println(msg)
	fmt.Print("Hello WOrld")
}
