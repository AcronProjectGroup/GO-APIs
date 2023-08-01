package store

// The Product struct defines Name and Category fields,
// which are exported,
// and a price field that is not exported.
type Product struct {
	Name, Category string
	price          float64
}

// Defining a Constructor:
func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}



// There is also a method named Price, which accepts a float64 parameter and uses it with the
// price field to calculate a tax-inclusive price.
func (p *Product) Price(taxRate float64) float64 {
	return p.price + (p.price * taxRate)
}



