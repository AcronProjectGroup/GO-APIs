package main

import (
	"fmt"
	"time"
)

var channel1 = make(chan string)
var channel2 = make(chan string)

func task1() {
	time.Sleep(1 * time.Second)
	channel1 <- "Channel1 <- ONE "
}

func task2() {
	time.Sleep(2 * time.Second)
	channel2 <- "Channel2 <- TWO "
}

func main() {
	go task1()
	go task2()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <- channel1:
			fmt.Println("received", msg1)
			fmt.Println(channel1)
		case msg2 := <- channel2:
			fmt.Println("received", msg2)
			fmt.Println(channel2)
		
		}
	}
}
