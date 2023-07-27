package main

type Product struct {
	name, category string
	price          float64
}

func (p Product) getName() string {
	return p.name
}

func (p Product) getCost(_ bool) float64 {
	return p.price
}


// You can force the use of references by specifying pointer receivers when implementing the interface
// methods
/*
func (p *Product) getName2() string {
	return p.name
}
func (p *Product) getCost2(_ bool) float64 {
	return p.price
}
*/
/*
	cannot use product (variable of type Product) as Expense value in variable declaration:
	Product does not implement Expense (missing method getCost)compilerInvalidIfaceAssign
*/
