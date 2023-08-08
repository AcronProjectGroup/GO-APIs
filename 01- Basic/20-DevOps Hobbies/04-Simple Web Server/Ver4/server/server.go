package server

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type server struct {
	counter int
	mutex   *sync.Mutex
}

func New() *server {
	return &server{
		counter: 0,
		mutex: &sync.Mutex{},
	}
}

func (s *server) Server(port int) {
	http.HandleFunc("/hello", s.hello)
	http.Handle("/sample", &sample{})
	http.HandleFunc("/bmi", s.bmi)
	http.HandleFunc("/counter", s.incrementCounter)

	addr := fmt.Sprintf(":%d", port)
	fmt.Println("Listening on:", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

type sample struct{}

func (s *sample) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sample handler.\n")
}
