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

var categoryMaxPrices = map[string]float64{
	"Watersports": 250 + (250 * defaultTaxRate),
	"Soccer":      150 + (150 * defaultTaxRate),
	"Chess":       50 + (50 * defaultTaxRate),
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

func (taxRate *taxRate) calcTax(product *Product) float64 {
	if product.price > taxRate.threshold {
		return product.price + (product.price * taxRate.rate)
	}
	return product.price
}
