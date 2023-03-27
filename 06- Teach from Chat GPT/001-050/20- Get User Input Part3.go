// Get User Input Part3

// Notion Source : https://sinalalenakhsh.notion.site/9-every-ways-get-user-input-in-Golang-ac2c758f26b04902b0c706194dc77c68

// https://go.dev/doc/

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("Enter your favorite color: ")

    scanner.Scan()

    color := scanner.Text()

    fmt.Printf("Your favorite color is %s.\n", color)
}

//--------------------------------------------
//--------------------------------------------
//--------------------------------------------
//--------------------------------------------
//--------------------------------------------

package main

import (
    "fmt"
    "os"
)

func main() {
    var username string

    fmt.Print("Enter your username: ")

    fmt.Fscanf(os.Stdin, "%s", &username)

    fmt.Printf("Your username is %s.\n", username)
}