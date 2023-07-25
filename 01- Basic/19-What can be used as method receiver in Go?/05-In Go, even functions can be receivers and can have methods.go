// In Go, even functions can be receivers and can have methods. 
// This can be used to enhance function with simple decorators. 

package main

import (
    "fmt"
    "strings"
)

type F func(string) string

func (f F) upperCase(s string) string {
    return strings.ToUpper(f(s))
}

func (f F) lowerCase(s string) string {
    return strings.ToLower(f(s))
}

func (f F) mixCase(s string) string {
    return f.lowerCase(s) + f.upperCase(s)
}

func doubleString(a string) string {
    return a + a
}

func main() {
    f := F(doubleString)

    fmt.Println(f("a"))

    fmt.Println(f.upperCase("a"))

    fmt.Println(f.lowerCase("a"))

    fmt.Println(f.mixCase("a"))
}

