package main

import (
    "fmt"
)

func main() {
    var num1, num2, num3, num4 float64

    fmt.Print("")
    _, err := fmt.Scan(&num1, &num2, &num3, &num4)

    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }

    fmt.Println( (num1 + num2 + num3 + num4) / 4 )
}
