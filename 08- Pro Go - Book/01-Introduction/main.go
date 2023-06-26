package main

import "fmt"

type Rsvp struct {
	Name, Email, Phone string
	WillAttend         bool
}

func main() {
	fmt.Println("TODO: add some features")

	var theMake = make([]*Rsvp, 0)

	theMake = append(theMake, &Rsvp{
		Name:       "Sina",
		Email:      "Email",
		Phone:      "123",
		WillAttend: true,
	})

	// Accessing the fields
	if len(theMake) > 0 {
		// Accessing the first element of the slice
		firstRsvp := theMake[0]

		// Accessing the Name field
		name := firstRsvp.Name
		fmt.Println("Name:", name)

		// Accessing the Email field
		email := firstRsvp.Email
		fmt.Println("Email:", email)

		fmt.Println(&theMake)
		fmt.Println(&firstRsvp)
		fmt.Println(firstRsvp)
	}

}
