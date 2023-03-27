// Get User Input Part2

// Notion Source : https://sinalalenakhsh.notion.site/9-every-ways-get-user-input-in-Golang-ac2c758f26b04902b0c706194dc77c68

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter your age: ")

    age, _ := reader.ReadString('\n')

	age = strings.TrimSuffix(age, "\n")
	
    fmt.Printf("You are %s years old.\n", age)


}