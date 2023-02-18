package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}

/*
اگه عدد منفی به تابع داده بشه تابع دوباره خودش رو فراخوانی میکنه ولی این بار ورودی رو منفی در منفی میکنه یعنی مثبت میشه
برای ورودی تابع ، برای عملیات ریاضی اول تبدیل به مثبت میشه 

Output:
1.4142135623730951     2i

*/ 