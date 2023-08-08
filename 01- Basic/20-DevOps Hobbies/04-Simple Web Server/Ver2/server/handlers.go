package server

import (
	"fmt"
	"net/http"
)


func (handler *server) hello(w http.ResponseWriter, Request *http.Request) {
	fmt.Fprintf(w, "Hello Acron Projections !!!\n")
}


