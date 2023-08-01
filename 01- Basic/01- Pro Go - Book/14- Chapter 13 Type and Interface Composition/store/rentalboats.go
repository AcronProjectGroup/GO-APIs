package store

type RentalBoat struct {
	*Boat
	IncludeCrew bool
}

// The RentalBoat type is composed using the *Boat type, 
// which is, in turn, composed using the *Product
// type, forming a chain.
func NewRentalBoat(name string, price float64, capacity int, motorized, crewed bool) *RentalBoat {
	return &RentalBoat{NewBoat(name, price, capacity, motorized), crewed}
}
