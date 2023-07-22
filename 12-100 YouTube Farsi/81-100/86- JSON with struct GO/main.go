/*

Step 1 − This program imports “encoding/json” and “fmt” as necessary packages

Step 2 − Create two structs named Person and Employee where these structs have fields Name, Age, and Address, with corresponding json tags

Step 3 − Create a main function

Step 4 − In the main, create an object of Person struct named person and initialize its fields with some values

Step 5 − In this step, create an object of the Employee struct named employee and initialize its fields with some values

Step 6 − Then, use json.Marshal() method to convert the person and employee structs to JSON format

Step 7 − Store the output obtained in personJSON and employeeJSON variables

Step 8 − Finally, print the JSON representation of the person and employee structs using fmt.Println().

*/

package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:”address”`
}

type Employee struct {
	Name    string `json:"emp_name"`
	Age     int    `json:"emp_age"`
	Address string `json:"emp_address"`
}

func main() {
	person := Person{
		Name:    "Karan",
		Age:     24,
		Address: "Hauz Khas",
	}

	employee := Employee{
		Name:    "Tarun",
		Age:     26,
		Address: "Saket",
	}

	personJSON, _ := json.Marshal(person)
	employeeJSON, _ := json.Marshal(employee)

	fmt.Println("Person JSON:", string(personJSON))
	fmt.Println("Employee JSON:", string(employeeJSON))
}



/*

Output:

	Person JSON: {"name":"Karan","age":24,"address":"Hauz Khas"}
	Employee JSON: {"emp_name":"Tarun","emp_age":26,"emp_address":"Saket"}

*/