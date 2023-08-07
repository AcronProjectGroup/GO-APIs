package main

import "fmt"

type Person struct {
	Name string
}

type UserName struct {
	Person
	userName string
}

func (Usr UserName) UserNameSayHello() {
	fmt.Println(Usr.userName, "say Hello to you...")
}



func main()  {

	Sina := UserName{
		Person: Person{
			Name: "SINA",
		},
		userName: "sinalalehbakhsh",
	}

	name := Sina.Name
	fmt.Println(name)

	Sina.UserNameSayHello()
}



