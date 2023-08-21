package main

import (  
	"fmt"
	"net"
)


func handleRequest(conn net.Conn) {
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {  
		fmt.Println("Error reading:", err.Error())  
		conn.Close()
		return
	}

	message := string(buffer)
	fmt.Println("Received message:", message)
	// Process the received message or perform any necessary actions
	response := "Hello, client!"
	conn.Write([]byte(response))
	conn.Close()

}



func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {  
		fmt.Println("Error listening:", err.Error())  
		return
		}
	defer listener.Close()
	fmt.Println("Server listening on localhost:8080")
	for {  
		conn, err := listener.Accept()  
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
			}  
		go handleRequest(conn)
		
	}

}


