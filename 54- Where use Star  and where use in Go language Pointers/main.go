// Where use Star * and where use & in Go language Pointers

package main



func main(){
	number := 123
	print("number= ",number, "\n")

	anotherNumber := &number
	print("anotherNumber= ",anotherNumber, "\n")
	
	*anotherNumber = 444
	print("anotherNumber= ",anotherNumber, "\n")

	// Where use Star? *
	print("*anotherNumber= ",*anotherNumber, "\n")

	print("Change Original 'number'--> ",number, "\n")


}