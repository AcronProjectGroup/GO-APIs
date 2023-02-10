package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	Clear_123()
	fmt.Println("Hello Sina ... ")
	animation()
}

func Clear_123() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func animation() {
	for i := 0; i <= 10; i++ {
		fmt.Print("*")                        
		time.Sleep(100 * time.Millisecond)
		Clear_123()
		fmt.Print("**")
		time.Sleep(100 * time.Millisecond)
		Clear_123()
		fmt.Print("***")
		time.Sleep(100 * time.Millisecond)
		Clear_123()
		fmt.Print("****")
		time.Sleep(100 * time.Millisecond)
		Clear_123()
	}
}