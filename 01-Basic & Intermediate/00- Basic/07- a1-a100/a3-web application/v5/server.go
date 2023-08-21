package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/pprof"
	"sync"
	"time"
)

// Server describe HTTP API server
type Server struct {
	http http.Server
	mux *http.ServeMux
}

const (
	httpAPITimeout = time.Second * 3
	shutdownTimeout = time.Second * 10
)

func New(port int) *Server {
	s := &Server{}
	mux := http.NewServeMux()

	mux.Handle("/debug/pprof/", http.HandleFunc(pprof.Index))
	mux.Handle("debug/pprof/cmdline", http.HandleFunc(pprof.Cmdline()))
	
	s.mux = mux
}
















