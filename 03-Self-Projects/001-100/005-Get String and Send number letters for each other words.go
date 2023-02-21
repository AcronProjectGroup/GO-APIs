// // Source : A Tour of Go 


// /* first take this :
// in Terminal:
// 	go mod init nameProject
// than: 
// 	go get golang.org/x/tour/wc 

// than: copy this in package main and use it

// */

// package main

// import (
// 	// "fmt"
// 	"strings"
// 	"golang.org/x/tour/wc"
// )

// func main() {
// 	// myMap := WordCount("sina lale bakhsh")
// 	// fmt.Println(myMap)
// 	wc.Test(WordCount)

// }



// // package main

// func WordCount(s string) map[string]int {
// 	myMap := map[string]int{}

// 	vari := strings.Fields(s)
// 	for i := range vari {
// 		myMap[vari[i]] = len(vari[i])  
		
// 	}
	
// 	return myMap
// }

// // func main() {
// // 	wc.Test(WordCount)
// // }

