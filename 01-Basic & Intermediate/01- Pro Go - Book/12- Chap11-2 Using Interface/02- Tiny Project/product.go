package main

type Product1 struct {
	name, category string
	price          float64
}

func (p Product1) getName1() string {
	return p.name
}

func (p Product1) getCost1(_ bool) float64 {
	return p.price
}


// You can force the use of references by specifying pointer receivers when implementing the interface
// methods
/*
func (p *Product1) getName2() string {
	return p.name
}
func (p *Product1) getCost2(_ bool) float64 {
	return p.price
}
*/
/*
	cannot use product (variable of type Product1) as Expense value in variable declaration:
	Product1 does not implement Expense (missing method getCost)compilerInvalidIfaceAssign
*/
