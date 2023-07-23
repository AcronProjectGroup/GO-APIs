// Installation of pprof: Your First Step Towards Optimized Go (Golang) Code

// https://voskan.host/2023/07/22/profiling-in-golang-with-pprof/

/*

Having understood the importance of profiling in Go,
our next step involves installing pprof,
a pivotal tool in the Go profiling ecosystem.

pprof is a package incorporated within Go’s standard library,
offering an array of profiling options to developers.

*/

// How to Install pprof in Go
// As pprof is part of Go’s standard library, there’s no separate installation process.

// After importing pprof, you can use its functionalities by starting an HTTP server.

package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()
}
