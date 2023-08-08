// https://sinalalenakhsh.notion.site/6-Struct-in-Golang-91c9d150796c4a20acf6bd69ed03c798

package main

import (
	"fmt"
)

type Car struct {
    make    string
    model   string
    year    int
    color   string
    mileage float64
    price   float64
    owner   string
}

func main() {
    myCar := Car{
        make:    "Toyota",
        model:   "Camry",
        year:    2021,
        color:   "Red",
        mileage: 12000,
        price:   25000,
        owner:   "John Smith",
    }
    fmt.Printf("My car is a %v %v %v with %v miles on it. It is %v and cost me $%.2f. It is owned by %v.\n", myCar.year, myCar.make, myCar.model, myCar.mileage, myCar.color, myCar.price, myCar.owner)
}

// ---------------------------------------------------


func main() {
	unumaq2500H := Car{
		makeBy:  "Mercedess",
		model:   "Unimaq",
		mileage: 10.43,
	}

	fmt.Printf("my Car make By %s, and model is %v, and mileage is %.1f Kilometer\n", unumaq2500H.makeBy, unumaq2500H.model, unumaq2500H.mileage)

	benzeCLS500 := Car{
		makeBy:  "Mercedess",
		model:   "CLS500",
		mileage: 120.476,
	}
	fmt.Printf("my Car make By %s, and model is %v, and mileage is %.1f Kilometer\n", benzeCLS500.makeBy, benzeCLS500.model, benzeCLS500.mileage)

}

type Car struct {
	makeBy  string
	model   string
	mileage float64
}
