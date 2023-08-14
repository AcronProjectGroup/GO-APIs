package main

import (
	"fmt"
	"time"
)

func Delay_4_Second(when_Channel_finished_send_done chan bool ) {
	for {
		fmt.Println("1")
		time.Sleep(time.Second)
		fmt.Println("2")
		time.Sleep(time.Second)
		break
	}
	// when_Channel_finished_send_done <- false
	close(when_Channel_finished_send_done)
}

func main()  {

		myChannel_is_finished := make(chan bool, 1) 
	
		fmt.Println("01- Avalin Statment!!!")
		
		go Delay_4_Second(myChannel_is_finished)
		
		fmt.Println("Aakharin Statment.........")
		

		for {
			if errorDetails, closeChannel := <- myChannel_is_finished; errorDetails{
				fmt.Println(errorDetails)
			} else {
				fmt.Println("Channel is Closed")
				fmt.Println(closeChannel)
				break
			}
		}
	

}



