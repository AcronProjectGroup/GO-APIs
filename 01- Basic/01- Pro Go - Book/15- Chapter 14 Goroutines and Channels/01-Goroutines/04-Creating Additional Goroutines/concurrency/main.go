/*

This statement tells the runtime to execute the statements 
in the TotalPrice method using a new goroutine. The runtime 
doesn’t wait for the goroutine to execute the method and 
immediately moves onto the next statement. This is the 
entire point of goroutines because the TotalPrice method 
will be invoked asynchronously, meaning that its statements 
are being evaluated by one goroutine at the same time that 
the original goroutine is executing the statements in the 
main function. But, as I explained earlier, the program is 
terminated when the main goroutine has executed all the 
statements in the main function. The result is that 
the program terminates before the goroutines are created 
to execute the TotalPrice method complete, which is why 
there are no subtotals. I explain how to address this 
problem as I introduce additional features, but, for the 
moment, all I need to do is prevent the program from 
ending long enough for the goroutines to complete



*/

package main

import "fmt"

func main() {
	fmt.Println("main function started")
	CalcStoreTotal(Products)
	fmt.Println("main function complete")
}




/*

go run main.go 
Ouput:
	# command-line-arguments
	./main.go:7:2: undefined: CalcStoreTotal
	./main.go:7:17: undefined: Products

go run .
Output:
	main function started
	Watersports subtotal: $328.95
	Soccer subtotal: $79554.45
	Chess subtotal: $1291.00
	Total: $81174.40
	main function complete

*/