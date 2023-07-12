package main

import (
	"fmt"
	"math/rand"
	"time"

)



func main()  {

	rangeNumber := 6

	nesf123 := rangeNumber / 2
	rand.Seed(time.Now().UnixNano())
	nesf123 += 1
    fmt.Println(rand.Intn(rangeNumber - nesf123 +1) + nesf123)

}











