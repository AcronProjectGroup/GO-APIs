package main

func main() {
	allNumbers := []int{    
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}

	comparator_Number := 0
	for _, value := range allNumbers {
		if comparator_Number < value {
			comparator_Number = value
		}
	}

	println(comparator_Number)
}