package main

import "fmt"

func main() {

    messages := make(chan string, 3)

    messages <- "buffered"
    messages <- "channel"
    messages <- "channel1"

    fmt.Println(<-messages)
    fmt.Println(<-messages)
    fmt.Println(<-messages)
}

/*

By default channels are unbuffered, 
meaning that they will only accept sends 
(chan <-) if there is a corresponding receive 
(<- chan) ready to receive the sent value. 
Buffered channels accept a limited 
number of values without a corresponding 
receiver for those values.

*/