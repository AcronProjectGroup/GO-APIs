// https://chat.openai.com/share/231c32ba-c375-4e6b-aa9d-fd341fce3a1d

package main

import (
	// "crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

func handleClient(client net.Conn, remoteAddr string) {
	remote, err := net.Dial("tcp", remoteAddr)
	if err != nil {
		log.Println("Failed to connect to remote:", err)
		return
	}
	defer remote.Close()

	go io.Copy(client, remote)
	io.Copy(remote, client)
}

func handleRequest(w http.ResponseWriter, r *http.Request, remoteAddr string) {
	client, _, err := w.(http.Hijacker).Hijack()
	if err != nil {
		log.Println("Failed to hijack connection:", err)
		return
	}
	defer client.Close()

	handleClient(client, remoteAddr)
}

func main() {
	localAddr := "0.0.0.0:8000"    // Local IP and port to listen on
	remoteAddr := "example.com:80" // Destination IP and port
	certFile := "cert.pem"         // Path to SSL certificate file
	keyFile := "key.pem"           // Path to SSL private key file

	listener, err := net.Listen("tcp", localAddr)
	if err != nil {
		log.Fatal("Failed to start listener:", err)
	}
	defer listener.Close()

	fmt.Println("Proxy is listening on", localAddr)
	fmt.Println("Forwarding packets to", remoteAddr)

	// Start web server with SSL/TLS support
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			handleRequest(w, r, remoteAddr)
		})

		err := http.ListenAndServeTLS(":8080", certFile, keyFile, nil)
		if err != nil {
			log.Fatal("Failed to start web server:", err)
		}
	}()

	for {
		client, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept client connection:", err)
			continue
		}

		go handleClient(client, remoteAddr)
	}
}



/*

In this updated version, we added two variables certFile and keyFile 
to specify the paths to the SSL certificate file and private key file respectively.

To obtain an SSL certificate, you can either generate a self-signed certificate 
for testing purposes or acquire a certificate from a trusted certificate authority (CA) for production use.

To generate a self-signed certificate, you can use tools like OpenSSL. 
Here are the steps to generate a self-signed certificate:

Generate a private key file:
	openssl genrsa -out key.pem 2048

Generate a certificate signing request (CSR) file:
	openssl req -new -key key.pem -out csr.pem

Generate a self-signed certificate using the private key and CSR:
    openssl x509 -req -days 365 -in csr.pem -signkey key.pem -out cert.pem

Make sure to replace the certFile and keyFile variables 
in the Go program with the actual paths to the generated certificate and key files.

Once you have the certificate and key files, you can run the program, 
and the web server will listen on port 8080 using HTTPS with the provided SSL certificate.

*/