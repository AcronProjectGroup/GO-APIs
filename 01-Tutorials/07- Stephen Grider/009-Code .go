//Sure! Here's an example that demonstrates the memory efficiency of using slices in Go:

// Create an array of 10 integers
myArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// Create a slice that references the first 5 elements of the array
mySlice := myArray[:5]

// Print the contents of the slice and array
fmt.Println(mySlice) // Output: [1 2 3 4 5]
fmt.Println(myArray) // Output: [1 2 3 4 5 6 7 8 9 10]

// Print the length and capacity of the slice
fmt.Println(len(mySlice)) // Output: 5
fmt.Println(cap(mySlice)) // Output: 10
