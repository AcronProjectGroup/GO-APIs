package main

import (
	"fmt"
)

func main() {
	var x, y int = 3, 4
	var f float64 = float64(x*x + y*y)
	var z uint = uint(f)
	fmt.Println(x, y, z)
}



/*
اگه زمان تغییر نوع سمت راست مساوی قبل از پرانتز دقیقا نوعش رو مشخص کنی درست انجام می شه
var f float64 = float64(x*x + y*y)      => Can Conversion to float64

اگه سمت راست  دقیقا مشخص نکنی + سمت چپ مساوی مشخص کنی ارور می ده
var f float64 = (x*x + y*y)      		=> ERROR

*/ 