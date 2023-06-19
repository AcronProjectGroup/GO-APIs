package main

import (
	"io/ioutil"
	"net/http"
)

func main()  {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
		b, _ := ioutil.ReadFile("cover-YouTube.jpg")
		w.Header().Add("content-type", "image/jpg")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	})

	http.ListenAndServe(":8888", nil)
	
}