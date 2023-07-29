package store

type Product struct {
	Name, Category string
	price          float64
}

func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}
func (p *Product) Price() float64 {
	return p.price
}
func (p *Product) SetPrice(newPrice float64) {
	p.price += newPrice 
}


/*

The name specified by the package statement should match the name of the folder
in which the code files are created,
which is store in this case.



The Product type has some important differences from similar types defined
in earlier chapters,
as I explain in the sections that follow.
*/
