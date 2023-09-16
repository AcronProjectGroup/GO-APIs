package main

import (
	"fmt"
	"time"
)

// func Delay_4_Second() {
// 	fmt.Println("1")
// 	time.Sleep(time.Second)
// 	fmt.Println("2")
// 	time.Sleep(time.Second)
// 	fmt.Println("3")
// 	time.Sleep(time.Second)
// 	fmt.Println("4")
// 	time.Sleep(time.Second)
// }

func Delay_4_Second(when_Channel_finished_send_done chan bool ) {
	fmt.Println("1")
	time.Sleep(time.Second)
	fmt.Println("2")
	time.Sleep(time.Second)
	fmt.Println("3")
	time.Sleep(time.Second)
	fmt.Println("4")
	time.Sleep(time.Second)
	when_Channel_finished_send_done <- true
}


func main()  {

	// 1-
		// fmt.Println("01- Avalin Statment!!!")
		// fmt.Println("Aakharin Statment.........")

		/* Output
			01- Avalin Statment!!!
			Aakharin Statment.........
		*/
		

	// // 2-
		// fmt.Println("01- Avalin Statment!!!")
		// Delay_4_Second()
		// fmt.Println("Aakharin Statment.........")
	
		/* Output:
			01- Avalin Statment!!!
			1
			2
			3
			4
			Aakharin Statment.........
		*/

	// 3-
		// fmt.Println("01- Avalin Statment!!!")
		// go Delay_4_Second()
		// fmt.Println("Aakharin Statment.........")
	
		/* Output:
			01- Avalin Statment!!!
			Aakharin Statment.........
		*/


	// 4-
		myChannel_is_finished := make(chan bool, 1) 
	
		fmt.Println("01- Avalin Statment!!!")
		
		go Delay_4_Second(myChannel_is_finished)
		
		fmt.Println("Aakharin Statment.........")
		
		<- myChannel_is_finished
	
		/* Output: 
			01- Avalin Statment!!!
			Aakharin Statment.........
			1
			2
			3
			4
		*/

}



