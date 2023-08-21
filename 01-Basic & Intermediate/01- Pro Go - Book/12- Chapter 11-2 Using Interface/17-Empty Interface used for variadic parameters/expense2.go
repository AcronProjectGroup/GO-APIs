package main

type ProductZ1 struct {
	name, category string
	price          float64
}

func (p ProductZ1) getName() string {
	return p.name
}

func (p ProductZ1) getCost(_ bool) float64 {
	return p.price
}
