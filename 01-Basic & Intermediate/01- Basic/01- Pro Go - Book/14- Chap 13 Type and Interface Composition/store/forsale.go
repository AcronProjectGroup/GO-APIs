package store

type ItemForSale interface {
	Price(taxRate float64) float64
}

/* Explanation:

Changing the map so that it uses 
the interface allows me to store Product and Boat values. 
The Product type conforms to the ItemForSale interface directly 
because there is a Price method that matches the signature specified 
by the interface and that has a *Product receiver.

There is no Price method that takes a *Boat receiver, 
but Go takes into account the Price method
promoted from the Boat type’s embedded field, 
which it uses to satisfy the interface requirements.


*/
