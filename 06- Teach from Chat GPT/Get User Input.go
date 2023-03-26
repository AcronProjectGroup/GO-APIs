// https://sinalalenakhsh.notion.site/5-Every-ways-get-user-input-in-Golang-ac2c758f26b04902b0c706194dc77c68



package main

import "fmt"

func main() {
    var name string

	for {
		fmt.Println("Enter your name: ")
		fmt.Print("-> ")
		fmt.Scan(&name)

		if name == "finish" {
			break
		}

		fmt.Printf("Hello, %s!\n", name)
	}


}

