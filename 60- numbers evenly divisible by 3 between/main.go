package main

import "time"

func main() {
	for i := 0; i < 100; i++ {
		if i % 3 == 0 {
			print(i, "\n")
			time.Sleep(time.Second * 1)
		}
	}
}