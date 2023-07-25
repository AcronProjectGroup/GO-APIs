package main

import "fmt"

type Person struct {
	Age int
	wight int
}

func (person Person) sumFields(k int) int {
	return person.Age + person.wight + k
}

func main() {
	person := Person{Age: 2, wight: 3}

	fmt.Println(person.sumFields(10))
}

/*



Question: 
	Can a function have multiple receiver arguments ?

Answer: 
	No it can’t. If you write func (a A, b B) f() {} 
	it will produce compile error method has multiple receivers. 
	You need to treat function receivers as the way to turn it into method of given type. 
	Just like in object oriented programming. 
	So to simplify, 
	our receiver can be compared to this or self keyword from object oriented languages.




*/