package main

import (
	"fmt"
	"time"
)

type Car struct {
	Name string
}

func PrintInAnotherTimeLine(Channel <- chan Car) {
	fromChannel := <-Channel
	fmt.Println(fromChannel.Name)
}

func main()  {
	var car Car = Car{
		Name: "Benze",
	}

	aChannel := make(chan Car, 1)
	aChannel <- car

	go PrintInAnotherTimeLine(aChannel)

	time.Sleep(time.Second)
	
}