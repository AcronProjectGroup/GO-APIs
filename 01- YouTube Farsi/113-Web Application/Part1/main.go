package main
import (
	"fmt"
	"net/http"
)
func handlerFunction(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "<h1>Welcome to AcronProject Site</h1>")
}
func main() {
	http.HandleFunc("/", handlerFunction)
	fmt.Println("Starting the Web Server on http://locanhost:3000")
	http.ListenAndServe(":3000", nil)
}
