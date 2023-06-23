// https://itnext.io/network-programming-in-go-389cd89ab874

package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	"golang.org/x/sync/semaphore"
)

func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	for {
		// Read the incoming connection into the buffer.
		len, err := conn.Read(buf)
		if err != nil {
			log.Printf("Error reading: %v", err)
			break
		}
		// Log the remote client's IP and port.
		log.Printf("Received connection from: %s", conn.RemoteAddr().String())

		// Log the time of the connection.
		log.Printf("Time of connection: %s", time.Now().Format(time.RFC3339))

		// Add to the total bytes received and log the new total.
		totalBytes += len
		log.Printf("Received %d bytes from the client. Total bytes received: %d", len, totalBytes)

		// Log the incoming message.
		log.Printf("Received message: %s", string(buf[:len]))
	}
	// Close the connection when you're done with it.
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error closing connection: %v", err)
		}
	}()
}



func main() {
	// Create a semaphore with a maximum weight of maxConcurrentConnections.
	sem = semaphore.NewWeighted(maxConcurrentConnections)

	// Create a log file.
	logFile, err := os.OpenFile("network.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer logFile.Close()

	// Set the output of the default logger to the log file.
	log.SetOutput(logFile)

	// Listen for incoming connections.
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatalf("Error listening: %v", err)
	}
	// Close the listener when the application closes.
	defer func() {
		if err = l.Close(); err != nil {
			log.Printf("Error closing listener: %v", err)
		}
	}()
	log.Println("Listening on localhost:2000")
	// ...

	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Error accepting: %v", err)
			continue
		}
		// Set KeepAlive for the connection.
		if tcpConn, ok := conn.(*net.TCPConn); ok {
			tcpConn.SetKeepAlive(true)
			tcpConn.SetKeepAlivePeriod(5 * time.Minute)
		}

		// Try to acquire the semaphore. If maxConcurrentConnections goroutines are already running, this will block
		// until one of them finishes and releases the semaphore.
		if err := sem.Acquire(context.Background(), 1); err != nil {
			log.Printf("Failed to acquire semaphore: %v", err)
			continue
		}

		// Handle connections in a new goroutine.
		go func() {
			// Release the semaphore when the goroutine finishes.
			defer sem.Release(1)
			handleRequest(conn)
		}()
	}

}
