// https://gobyexample.com/http-server


package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {

    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)

    http.ListenAndServe(":8090", nil)
}



/*

A fundamental concept in net/http servers is handlers. 
A handler is an object implementing the http.Handler interface. 
A common way to write a handler is by using the http.HandlerFunc 
adapter on functions with the appropriate signature.



*/