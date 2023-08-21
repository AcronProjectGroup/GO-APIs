package data

import (
	"fmt"
	"github.com/fatih/color" 
)

func init() {
	color.Magenta("\ndata.go init function invoked")
}


func GetData() []string {
	fmt.Println(("data.go init function invoked"))
	return []string{"Kayak", "Lifejacket", "Paddle", "Soccer Ball"}
}



// func init() {
// 	fmt.Println(("data.go init function invoked"))
// }


// func init() {
// 	fmt.Println(("data.go init function invoked"))
// }


// func init() {
// 	fmt.Println(("data.go init function invoked"))
// }

