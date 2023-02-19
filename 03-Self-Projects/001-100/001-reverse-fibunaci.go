package main

import "fmt"

// fib returns a function that returns
// successive Fibonacci numbers.
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func run_fib() {
	f := fib()
	for i := 0; i<= 10-1; i++{
		fmt.Print(f(), " ")
		fmt.Print(f(), " ")
		fmt.Print(f(), " ")
		fmt.Print(f(), " ")
		fmt.Print(f(), " ")
		fmt.Println()
	}
}

func return_run_fib() []int{
	f := fib()
	listNumbers := []int{}
	for i := 0; i<= 50-1; i++{
		listNumbers = append(listNumbers, f())
	}
	return listNumbers
}

func reverse() {
	listNumber := return_run_fib()
	for i := range listNumber {
		defer fmt.Println(listNumber[i])
	}
}

func main() {
	run_fib()
	reverse()
}

/* Output

1 1 2 3 5
8 13 21 34 55
89 144 233 377 610
987 1597 2584 4181 6765
10946 17711 28657 46368 75025
121393 196418 317811 514229 832040
1346269 2178309 3524578 5702887 9227465
14930352 24157817 39088169 63245986 102334155
165580141 267914296 433494437 701408733 1134903170
1836311903 2971215073 4807526976 7778742049 12586269025
20365011074 32951280099 53316291173 86267571272 139583862445

*/