/*

The sleep period is specified using a set 
of numeric constants that represent intervals so that time.
Second represents one second and is multiplied 
by 5 to create a five-second period.

*/

package main

import (
	"fmt"
	"time"
)


func main() {
	fmt.Println("main function started")

	CalcStoreTotal(Products)

	time.Sleep(time.Second)
	fmt.Println("1")
	time.Sleep(time.Second)
	fmt.Println("2")
	time.Sleep(time.Second)
	fmt.Println("3")
	time.Sleep(time.Second)

	fmt.Println("main function complete")

}


/*


 */
