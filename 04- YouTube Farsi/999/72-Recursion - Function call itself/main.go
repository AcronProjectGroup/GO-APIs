// 72-Recursion - Function call itself

/*

a function is able to call itself

5!= 1 * 2 * 3 * 4 * 5 = 120

6!= 1 * 2 * 3 * 4 * 5 * 6 = 720

*/

package main

func main() {
	factorial(2)
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

/*


factorial(2)
• Is x == 0? No. (x is 2)
• Find the factorial of x – 1
	• Is x == 0? No. (x is 1)
	• Find the factorial of x – 1
		• Is x == 0? Yes, return 1.
	• return 1 * 1
• return 2 * 1

*/
