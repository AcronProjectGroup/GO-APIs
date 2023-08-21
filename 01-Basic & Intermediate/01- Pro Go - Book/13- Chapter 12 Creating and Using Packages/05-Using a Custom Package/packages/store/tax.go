/*

All the features defined in the tax.go file are unexported,
which means they can be used only within the store package.

Notice that the calcTax method can access the price field of the Product type
and that it does so
without having to refer to the type as store.Product because it is in the same package


*/

package store

const defaultTaxRate float64 = 0.2

const minThreshold = 10

// Using an Initialization Function
var categoryMaxPrices = map[string]float64{
	"Watersports": 250,
	"Soccer":      150,
	"Chess":       50,
}
/*
	var categoryMaxPrices = map[string]float64 {
		"Watersports": 250 + (250 * defaultTaxRate),
		"Soccer": 150 + (150 * defaultTaxRate),
		"Chess": 50 + (50 * defaultTaxRate),
	}
*/

// ----------------------------------------------------
// ----------------------------------------------------
// ----------------------------------------------------


/*
	The initialization function is called init,
	and it is defined without parameters and a result.
	The init function is called automatically
	and provides an opportunity to prepare the package for use.

	The init function is not a regular Go function 
	and cannot be invoked directly. And, unlike regular
	functions, a single file can define multiple init functions, 
	all of which will be executed.
	
*/
func init() {
	for category, price := range categoryMaxPrices {
		categoryMaxPrices[category] = price + (price * defaultTaxRate)
	}
}

type taxRate struct {
	rate, threshold float64
}

func newTaxRate(rate, threshold float64) *taxRate {
	if rate == 0 {
		rate = defaultTaxRate
	}
	if threshold < minThreshold {
		threshold = minThreshold
	}
	return &taxRate{rate, threshold}
}

func (taxRate *taxRate) calcTax(product *Product) (price float64) {
	if product.price > taxRate.threshold {
		price = product.price + (product.price * taxRate.rate)
	} else {
		price = product.price
	}
	if max, ok := categoryMaxPrices[product.Category]; ok && price > max {
		price = max
	}
	return
}
