package main

type Product2 struct {
	name, category string
	price          float64
}

func (p Product2) getName2() string {
	return p.name
}

func (p Product2) getCost2(_ bool) float64 {
	return p.price
}

