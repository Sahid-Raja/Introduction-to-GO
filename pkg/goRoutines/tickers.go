package goroutines

import (
	"fmt"
	"time"
)

func Tickers() {
	ticker := time.NewTicker(1000 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				{
					fmt.Println("So Program is Done")
				}
			case <-ticker.C:
				{
					fmt.Println(time.Now())
				}

			}
		}
	}()

	time.Sleep(time.Millisecond * 4000)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
