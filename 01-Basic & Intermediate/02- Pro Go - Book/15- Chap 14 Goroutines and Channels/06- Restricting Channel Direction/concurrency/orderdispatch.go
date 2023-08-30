package main

import (
	"fmt"
	"math/rand"
	"time"
)

type DispatchNotification struct {
	Customer string
	*Product
	Quantity int
}

var Customers = []string{"Alice", "Bob", "Charlie", "Dora"}

func DispatchOrders(channel chan DispatchNotification) {
	rand.Seed(time.Now().UTC().UnixNano())
	orderCount := rand.Intn(3) + 2
	fmt.Println("Order count:", orderCount)
	for i := 0; i < orderCount; i++ {
		channel <- DispatchNotification{
			Customer: Customers[rand.Intn(len(Customers)-1)],
			Quantity: rand.Intn(10),
			Product:  ProductList[rand.Intn(len(ProductList)-1)],
		}
		if i == 1 {
			notification := <-channel
			fmt.Println("Read:", notification.Customer)
		}
	}
	close(channel)
}

/* Output:

ubuntu@ubuntu-pass-123:~/Desktop/Repo/Test/Go Pro/concurrency$ go run .
Order count: 3
Read: Bob
Dispatch to Bob : 3 x Thinking Cap
Dispatch to Alice : 6 x Unsteady Chair
Channel has been closed

ubuntu@ubuntu-pass-123:~/Desktop/Repo/Test/Go Pro/concurrency$ go run .
Order count: 3
Read: Bob
Dispatch to Alice : 5 x Kayak
Dispatch to Charlie : 9 x Thinking Cap
Channel has been closed

ubuntu@ubuntu-pass-123:~/Desktop/Repo/Test/Go Pro/concurrency$ go run .
Order count: 4
Read: Alice
Dispatch to Bob : 1 x Soccer Ball
Dispatch to Bob : 2 x Unsteady Chair
Dispatch to Bob : 1 x Corner Flags
Channel has been closed

*/