// one .go file = go run main.go

// more than one .go file = go run main.go another.go andAnother.go

// more than one .go file = go run . = ALL .go files runs

package main

import (
	"fmt"
	"time"

	"github.com/jalaali/go-jalaali"
)

func main()  {
	
	Year, month, Day, err:= jalaali.ToJalaali(time.Now().Date())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println( Day , month  , Year)
}

