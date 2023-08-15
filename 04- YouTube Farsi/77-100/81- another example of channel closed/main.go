package main

import (
	"fmt"
	"time"
)

func Delay_4_Second(when_Channel_finished_send_done chan bool) {
	fmt.Println("1")
	time.Sleep(time.Second)
	fmt.Println("2")
	time.Sleep(time.Second)

	// The second mode 2-1
	// when_Channel_finished_send_done <- false

	// The third mode 3-1
	// when_Channel_finished_send_done <- false
	// when_Channel_finished_send_done <- true
	// close(when_Channel_finished_send_done)

	myTrue := false
	if myTrue {
		when_Channel_finished_send_done <- true
	} else if myTrue == false {
		close(when_Channel_finished_send_done)
	}
}

func main() {

	myChannel_is_finished := make(chan bool, 1)
	fmt.Println("Avalin Statment!!!")

	go Delay_4_Second(myChannel_is_finished)

	fmt.Println("Aakharin Statment.........")

	// The first mode for concurrency
	// time.Sleep(time.Microsecond * 2000000 )

	// The second mode 2-2
	// <-myChannel_is_finished
	// fmt.Println(<-myChannel_is_finished)
	// valueOfChannel := <-myChannel_is_finished
	// fmt.Println(valueOfChannel)

	// The third mode 3-2
	for {
		if details, open := <-myChannel_is_finished; open {
			fmt.Println("OPEN")
			fmt.Println(details) // details are values of channel
			break
		} else if !open {
			fmt.Println("CLOSE")
			fmt.Println(details) // details are values of channel
			break
		}
	}
}

