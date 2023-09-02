package main

import (
	"fmt"
	"strings"
)

func main() {
	username := " A l   i c e"
	
	trimmed := strings.TrimSpace(username)
	
	fmt.Println("Trimmed:", ">>"+trimmed+"<<")

}
