//Difference in behavior between slices and maps in Golang

// Source Notion:  

// stackoverflow.com:  https://stackoverflow.com/questions/52681793/difference-in-behavior-between-slices-and-maps

func main() {
    var s []int // len and cap are both 0
    var m map[int]int

    fmt.Println(s) // works... prints an empty slice
    fmt.Println(m) // works... prints an empty map

    s = append(s, 10) // returns a new slice, so underlying array gets allocated
    fmt.Println(s) // works... prints [10]

    m[10] = 10 // program crashes, with "assignment to entry in nil map"
    fmt.Println(m)
}


func main() {
    var mySlice []int 
    var myMap = make(map[int]int)
    mySlice = append(mySlice, 10,12,13,14,15) 
    for _, value := range mySlice {
        fmt.Println(value)
    }
    myMap = map[int]int{
        1: 1,
        2: 2,
        3: 3,
        4: 4,
        5: 5,
    }
    for _, value := range myMap {
        fmt.Println(value)
    }
}