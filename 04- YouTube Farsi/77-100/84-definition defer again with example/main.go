package main

import (
	"fmt"
	"time"
)

func trackTime(t time.Time) {
	elapsed := time.Since(t)
	fmt.Println("execution completed in:", elapsed)
}

func main()  {

	trackTime(time.Now())

	defer trackTime(time.Now())

	time.Sleep(time.Second * 2 )


}
