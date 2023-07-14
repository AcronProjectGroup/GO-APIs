package main

import (
	"bufio"
	"net"
)

func main() {
	// Step1-> Listen
	// Step2-> Accept
	// Step3-> Handle connection -> thread

	// Step1-> Listen
	streamServer, err := net.Listen("tcp", ":60123") // For closing: Ctrl + ]
	if err != nil {
		print(err)
		return
	}
	defer streamServer.Close()

	// Step2-> Accept
	for {
		connectionUser, err := streamServer.Accept()
		if err != nil {
			print(err)
			return
		}
		go handle(connectionUser)
	}
}

// Step3-> Handle connection -> thread
func handle(connectionUser net.Conn) {
	for {
		data, err := bufio.NewReader(connectionUser).ReadString('\n') // Enter Key on Keyboard
		if err != nil {
			println("GoodBye")
			connectionUser.Close()
			return
		}
		print(data)
	}
}
