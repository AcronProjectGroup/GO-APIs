package main

import "time"

func main() {
	for i := 0; i <= 100; i++ {
		time.Sleep(time.Second * 1)
		if i % 5 == 0 {
			println("HOB")
		} else {
			println(i)
		}
	}
}
















