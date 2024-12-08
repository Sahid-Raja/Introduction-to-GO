package goroutines

import (
	"fmt"
	"time"
)

func recieveOrder(orders chan<- string) {
	for i := 1; i <= 50; i++ {
		order := fmt.Sprintf("Order :#%v", i)
		fmt.Println("R ", order)
		orders <- order
		time.Sleep(300 * time.Millisecond)
	}
	close(orders)
}

func createOrder(orders <-chan string, packed chan<- string) {
	for order := range orders {
		fmt.Printf("P %v\n", order)
		packedOrder := order + " (packed) "
		packed <- packedOrder
		time.Sleep(400 * time.Millisecond)
	}
	close(packed)
}

func dispatchOrder(packed <-chan string) {
	for packedOrder := range packed {
		fmt.Println("D ", packedOrder)
		time.Sleep(500 * time.Millisecond)
	}
}

func ChannelSyncronization() {
	fmt.Println("Shop is Open Now!!!")
	orders := make(chan string, 2)
	packed := make(chan string, 2)

	go recieveOrder(orders)
	go createOrder(orders, packed)
	go dispatchOrder(packed)

	time.Sleep(30 * time.Second)
	fmt.Println("Shop is Closed now....")

}
