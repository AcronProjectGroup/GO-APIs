package server

import (
	"fmt"
	"log"
	"net/http"
)

type server struct {}


func New() *server {
	return &server{}
}


func (s *server) Server(port int) {
	http.HandleFunc("/hello", s.hello)


	addr := fmt.Sprintf(":%d", port)
	fmt.Println("Listening on:", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
