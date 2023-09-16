package main

import (
	"fmt"
	"strings"
)

func main() {
	userNames := map[string]string{
		"acronproject": "sinalale",
	}

	fmt.Println("Enter a name:")

	var name string
	_, err := fmt.Scanln(&name)
	if err != nil {
		panic(err)
	}

	name = strings.TrimSpace(name)

	if keyOfUserNames, ok := userNames[name]; ok {
		fmt.Printf("%q is the username of %q.\n", keyOfUserNames, name)
	}
	

	fmt.Printf("I don't have %s's username, What is it name?\n", name)

	var NewUserName string
	_, err = fmt.Scanln(&NewUserName)
	if err != nil {
		panic(err)
	}

	NewUserName = strings.TrimSpace(NewUserName)

	userNames[name] = NewUserName

	fmt.Println("Data updated.")
	fmt.Println("UserNames map:")
	fmt.Printf("%q\n",userNames)



}