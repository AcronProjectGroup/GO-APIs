/*

Rune:
	noun- سخن, سخن مرموز, نشان مرموز



Note the use of string(r) to convert the rune back to a string for comparison. 
This approach works, but it is less efficient because it involves creating 
a string for each rune comparison. 

Using []rune allows for direct rune-to-rune comparison, which is more efficient and simpler.

func isVowel(r rune) bool {
	vowels := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}
	for _, vowel := range vowels {
		if string(r) == vowel {
			return true
		}
	}
	return false
}


زمانی که قراره یک تک کاراکتر رو با یک تک کاراکتر مقایسه کنیم بهتره از روون استفاده کنیم
از استرینگ هم میشه استفاده کرد ولی اول استرینگ به روون تبدیل میشه بعدش مقایسه میشه
بعد دوباره به استرینگ تبدیل میشه و به ما نمایش داده میشه

*/



package main

import (
	"fmt"
)

func main() {
	str := "Hello, 世界!"
	filteredStr := ""
	for _, char := range str {
		if !isVowel(char) {
			filteredStr += string(char)
		}
	}
	fmt.Println(filteredStr)
}

func isVowel(r rune) bool {
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	for _, vowel := range vowels {
		if r == vowel {
			return true
		}
	}
	return false
}
