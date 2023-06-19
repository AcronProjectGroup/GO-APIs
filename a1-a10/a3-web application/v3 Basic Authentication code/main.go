/*

This is Basic Authentication code

*/

package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
)

type book struct {
	Isbn string `json:"isbn"`
	Name string `json:"name"`
}
var bookList []book = []book{}

func main()  {
	
	mux := http.DefaultServeMux
	mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
		b, _ := ioutil.ReadFile("cover-YouTube.jpg")
		w.Header().Add("content-type", "image/jpg")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	})

	
	mux.Handle("/book", http.HandlerFunc(handleBooks))

	http.ListenAndServe(":8888", nil)

}


func handleBooks(w http.ResponseWriter, r *http.Request){
	u, p, ok := r.BasicAuth()
	if !ok || u != "ADMIN" || p != "ADMIN" {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	switch r.Method {
	case "GET":
		b, _ := json.Marshal(bookList)
		w.Write(b)
		return
	case "POST":
		newBook := book{}
		b, _ := io.ReadAll(r.Body)
		json.Unmarshal(b, &newBook)
		for _, b := range bookList {
			if b.Isbn == newBook.Isbn {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
		
		bookList = append(bookList, newBook)
		w.WriteHeader(http.StatusAccepted)
		return

	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}


