// Defining a Constructor Function
package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}
func newProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}


type Person struct {
	name string
	sureName string
}



// A constructor function is responsible
// for creating struct values using values received through parameters
// Constructor functions are used to create struct values consistently.
// توابع سازنده برای ایجاد مقادیر ساختار به طور مداوم استفاده می شود.
// Constructor functions are usually named new or New followed 
// by the struct type so that the constructor function 
// for creating Product values is named newProduct.
// Constructor functions return struct pointers, 
// and the address operator is used directly with the literal struct syntax

// Constructor Function:
func newPerson(name, sureName string) *Person {
	return &Person{name, sureName}
}

// Modifying a Constructor:
	// meaning that I don’t have to find all the points 
	// in the code where Product values are created and modify them individually.
	// func newProduct(name, category string, price float64) *Product {
		// return &Product{name, category, price - 10}
	// }
	// ...


func main() {
	
	// important is that you remember to return a pointer to avoid the
	// struct value being duplicated when the function returns.
	// uses an array to store product data,
	// and you can see the use of pointers in the array type
	// ensuring that changes to the construction
	// process are reflected in all the struct values created by the function.
	products := [2]*Product{
		newProduct("Kayak", "Watersports", 275),
		newProduct("Hat", "Skiing", 42.50),
	}
	for _, p := range products {
		fmt.Println("Name:", p.name, "Category:", p.category, "Price", p.price)
	}



	persons := []*Person{
		newPerson("SINA", "LALEHBAKHSH"),
		newPerson("MINA", "JAMES"),
		newPerson("LINA", "FAINES"),
		newPerson("DINA", "LAINS"),
		newPerson("JINA", "VAINES"),
	}


	for index, person := range persons {
		fmt.Printf("Index-%d: %s-%s.\n", index, person.name, person.sureName)
	}
	
}

/* Output:

Name: Kayak Category: Watersports Price 275
Name: Hat Category: Skiing Price 42.5
Index-0: SINA-LALEHBAKHSH.
Index-1: MINA-JAMES.
Index-2: LINA-FAINES.
Index-3: DINA-LAINS.
Index-4: JINA-VAINES.

*/