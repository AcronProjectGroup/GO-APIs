package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	mx1 := http.DefaultServeMux
	mx1.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, _ := ioutil.ReadFile("Sina-Laleh.jpg")
		w.Header().Add("content-type", "image/jpg")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	})

	http.ListenAndServe(":8543", nil)
}

/* Version1

	mx1 := http.NewServeMux
	mx1.HandleFunc("/home",)
	mx1.HandleFunc("/about",)

	mx2 := http.DefaultServeMux
	mx2.HandleFunc("/sinalalebakhsh",)
	mx2.HandleFunc("/acronproject",)
	// In here everything rewrite to "mx2"
	mx2Plus := http.DefaultServeMux
	mx2Plus.HandleFunc("/sinalalebakhsh",)
	mx2Plus.HandleFunc("/acronproject",)

*/