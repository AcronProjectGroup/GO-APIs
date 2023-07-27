package main

type Product4 struct {
	name, category string
	price          float64
}

func (p Product4) getName2() string {
	return p.name
}

func (p Product4) getCost2(_ bool) float64 {
	return p.price
}

