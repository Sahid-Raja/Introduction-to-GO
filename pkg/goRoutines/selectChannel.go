package goroutines

import (
	"fmt"
	"time"
)

func addHelloWorld(ch1 chan<- string) {
	for i := 0; i < 2; i++ {
		ch1 <- "Hello World"
		time.Sleep(time.Millisecond * 500)
	}
	close(ch1)
}
func addHelloUniverse(ch2 chan<- string) {
	for i := 0; i < 2; i++ {
		ch2 <- "Hello Universe"
		time.Sleep(time.Millisecond * 2000)
	}
	close(ch2)
}
func SelectChannel() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go addHelloWorld(ch1)
	go addHelloUniverse(ch2)

	for i := 0; i < 10; i++ {
		select {
		case <-ch1:
			fmt.Println("Got Data From Channel 1")
			time.Sleep(time.Millisecond * 500)
		case <-ch2:
			fmt.Println("Got Data From Channel 2")
			time.Sleep(time.Millisecond * 500)
		default:
			fmt.Println("Nothing is Recieved")
			time.Sleep(time.Millisecond * 500)
		}
	}

}
