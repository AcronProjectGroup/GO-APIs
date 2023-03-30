/*
	30- create simple shell with Go lang like Bash

	Package os in Golang
	Important Source:      		https://sinalalenakhsh.notion.site/1-Chat-GPT-310ba8a7e074493e95d8c45bfe6aed86
	Current Notion Source: 		https://sinalalenakhsh.notion.site/14-Package-os-Part-1-2-3-4-2799ac556ef14fe294c8898413373dd1	

*/



package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(" *** in Go Program *** ")

	pwd , _  := reader.ReadString('\n')
	pwd = strings.TrimSuffix(pwd, "\n")


	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	
	if pwd == "pwd" {
		fmt.Println(dir)
	}

}
