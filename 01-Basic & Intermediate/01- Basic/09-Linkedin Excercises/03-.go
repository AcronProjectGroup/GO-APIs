package main

import (
	"fmt"
	"time"
)

func main()  {

	for _, v := range []string{"1", "2", "3"} {
		go fmt.Println(v)
	}
	time.Sleep(time.Second)


	for _, v := range []string{"1", "2", "3"} {
		go func() {
			fmt.Println(v)
		}()
	}
	time.Sleep(time.Second)

	
	
}