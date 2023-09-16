package main

import (
	"fmt"
	"errors"
)


func dividePlus(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("taghsim bar sefr nadaarim!!!")
	}
	return a / b, nil 
}

func main() {
	resultdivided, err := dividePlus(8, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultdivided)
	}

}
