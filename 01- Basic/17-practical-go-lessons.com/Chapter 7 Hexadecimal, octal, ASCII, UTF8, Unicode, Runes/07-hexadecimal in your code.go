// If you want to represent a number in hexadecimal in your code, add 0x before the numeral :

// hexadecimal-octal-ascii-utf8-unicode-runes/hex-number/main.go
package main

import "fmt"

func main() {
	n := 2548
	n2 := 0x9F4
	fmt.Printf("%X\n", n)
	fmt.Printf("%x\n", n2)
}

// Output :

// 9F4
// 9f4
