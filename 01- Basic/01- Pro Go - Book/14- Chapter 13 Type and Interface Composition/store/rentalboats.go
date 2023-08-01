package store

type Crew struct {
	Captain, FirstOfficer string
}

type RentalBoat struct {
	*Boat
	IncludeCrew bool
	*Crew
}

// The RentalBoat type is composed using the *Boat type,
// which is, in turn, composed using the *Product
// type, forming a chain.
// func NewRentalBoat(name string, price float64, capacity int, motorized, crewed bool) *RentalBoat {
// 	return &RentalBoat{NewBoat(name, price, capacity, motorized), crewed}
// }

func NewRentalBoat(
	name string, 
	price float64, 
	capacity int,
	motorized, crewed bool, 
	captain, firstOfficer string) *RentalBoat {
		return &RentalBoat{
			NewBoat(name, price, capacity, motorized), 
			crewed,
			&Crew{captain, firstOfficer},
		}
}

