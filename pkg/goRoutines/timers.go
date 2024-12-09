package goroutines

import (
	"fmt"
	"time"
)

func Timers() {
	timer1 := time.NewTimer(1 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 Fired", time.Now())

	timer2 := time.NewTimer(time.Microsecond)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 is Fired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped", time.Now())
	}
	time.Sleep(2 * time.Second)
}
