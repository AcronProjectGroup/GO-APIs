package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type book struct {
	Isbn string `json:"isbn"`
	Name string `json:"name"`
}

var bookList []book = []book{}

func main() {
	mx1 := http.DefaultServeMux
	mx1.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, _ := ioutil.ReadFile("Sina-Laleh.jpg")
		w.Header().Add("content-type", "image/jpg")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	})

	mx1.Handle("/homepage",http.HandlerFunc(Handle_HomePage))


	http.ListenAndServe(":8543", nil)
}

func Handle_HomePage(w http.ResponseWriter, r *http.Request){

	// if r.Method == "GET" {} else if r.Method == "POST" {}
	// Or we can write like this:
	switch r.Method {
	case "GET":
		b, _ := json.Marshal(bookList)
		w.Write(b)

	case "POST":
		newbook := book{}
		b, _ := io.ReadAll(r.Body)
		json.Unmarshal(b, &newbook)

		for _, b := range bookList {
			if b.Isbn == newbook.Isbn {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}

		bookList = append(bookList, newbook)
		w.WriteHeader(http.StatusAccepted)
		return

	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
}




























 







/* Version2
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

*/

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