// https://gobyexample.com/exit

package main

import (
    "fmt"
    "os"
)

func main() {

    defer fmt.Println("!")

    os.Exit(3)
}

/*

defers will not be run when using os.Exit, so this fmt.Println will never be called.

Note that unlike e.g. C, Go does not use an integer return value from main 
to indicate exit status. 
If you’d like to exit with a non-zero status you should use os.Exit.



*/