package main

import "fmt"

type rect struct {
    width, height int
}

func (r *rect) area() int {
    return r.width * r.height
}

func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
    sina := rect{width: 10, height: 5}

    fmt.Println("area: ", sina.area())
    fmt.Println("perim:", sina.perim())

    rp := &sina
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
}