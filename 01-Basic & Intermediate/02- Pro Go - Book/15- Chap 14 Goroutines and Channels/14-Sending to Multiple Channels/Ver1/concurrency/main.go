/*

If there are multiple channels available,
a select statement can be used to find a channel for which sending will not block

You can combine case statements with send and receive operations in the same select statement.
When the select statement is executed, the Go runtime builds a combined list of case statements that can be
executed without blocking and picks one at random, which can be either a send or a receive statement.


*/

package main

import (
	"fmt"
	"time"
)

func enumerateProducts(channel1, channel2 chan<- *Product) {
	for _, p := range ProductList {
		select {
		case channel1 <- p:
			fmt.Println("Send via channel 1")
		case channel2 <- p:
			fmt.Println("Send via channel 2")
		}
	}
	close(channel1)
	close(channel2)
}


func main() {
	c1 := make(chan *Product, 2)
	c2 := make(chan *Product, 2)
	go enumerateProducts(c1, c2)
	time.Sleep(time.Second)
	for p := range c1 {
		fmt.Println("Channel 1 received product:", p.Name)
	}
	for p := range c2 {
		fmt.Println("Channel 2 received product:", p.Name)
	}
}

/*

Output:

 go run .
Send via channel 2
Send via channel 2
Send via channel 1
Send via channel 1
Channel 1 received product: Soccer Ball
Channel 1 received product: Corner Flags
Channel 1 received product: Stadium
Send via channel 1
Send via channel 1
Send via channel 1
Send via channel 1
Channel 1 received product: Thinking Cap
Channel 1 received product: Unsteady Chair
Channel 1 received product: Bling-Bling King
Channel 2 received product: Kayak
Channel 2 received product: Lifejacket
ubuntu@ubuntu-pass-123:~/Desktop/Repo/Test/concurrency$ go run .
Send via channel 1
Send via channel 1
Send via channel 2
Send via channel 2
Channel 1 received product: Kayak
Channel 1 received product: Lifejacket
Channel 1 received product: Stadium
Send via channel 1
Send via channel 1
Send via channel 1
Send via channel 1
Channel 1 received product: Thinking Cap
Channel 1 received product: Unsteady Chair
Channel 1 received product: Bling-Bling King
Channel 2 received product: Soccer Ball
Channel 2 received product: Corner Flags
ubuntu@ubuntu-pass-123:~/Desktop/Repo/Test/concurrency$ go run .
Send via channel 2
Send via channel 2
Send via channel 1
Send via channel 1
Channel 1 received product: Soccer Ball
Channel 1 received product: Corner Flags
Channel 1 received product: Stadium
Send via channel 1
Send via channel 1
Send via channel 1
Send via channel 1
Channel 1 received product: Thinking Cap
Channel 1 received product: Unsteady Chair
Channel 1 received product: Bling-Bling King
Channel 2 received product: Kayak
Channel 2 received product: Lifejacket
ubuntu@ubuntu-pass-123:~/Desktop/Repo/Test/concurrency$ go run .
Send via channel 1
Send via channel 2
Send via channel 1
Send via channel 2
Channel 1 received product: Kayak
Channel 1 received product: Soccer Ball
Channel 1 received product: Stadium
Send via channel 1
Send via channel 1
Send via channel 1
Send via channel 1
Channel 1 received product: Thinking Cap
Channel 1 received product: Unsteady Chair
Channel 1 received product: Bling-Bling King
Channel 2 received product: Lifejacket
Channel 2 received product: Corner Flags

*/
