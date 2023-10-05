// Source : https://www.golangprograms.com/go-language/golang-maps.html


package main
 
import "fmt"
 
func main() {
	var employee = make(map[string]int)
	employee["Mark"] = 10
	employee["Sandy"] = 20
	fmt.Println(employee)
 
	employeeList := make(map[string]int)
	employeeList["Mark"] = 10
	employeeList["Sandy"] = 20
	fmt.Println(employeeList)

	var employee123  = map[string]int{"Sina": 29, "dddd": 13}
	employee123["Mark"] = 10
	employee123["Sandy"] = 20
	fmt.Println(employee123["Sina"])
}

/* Output

map[Mark:10 Sandy:20]
map[Mark:10 Sandy:20]
29

*/