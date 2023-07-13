package main

import (
	"fmt"
	"math/rand"
	"time"

)



func main()  {
	min := 0
	max := 3

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(max - min) + min

	fmt.Println(randomNumber)
}











