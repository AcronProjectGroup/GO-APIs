package main

import (
    "fmt"
    "os"
)

func main() {
    var x int = 42
    var s string = "hello"

    fileUs, err := os.Create("FileUs.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    defer fileUs.Close()
    fmt.Fprintf(fileUs, "x = %d, s = %s", x, s)

}