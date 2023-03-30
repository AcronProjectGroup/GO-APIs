/*
	Package os in Golang
	Important Source:      		https://sinalalenakhsh.notion.site/1-Chat-GPT-310ba8a7e074493e95d8c45bfe6aed86
	Current Notion Source: 		https://sinalalenakhsh.notion.site/14-Package-os-Part-1-2-3-4-2799ac556ef14fe294c8898413373dd1	

*/

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// 1. Get current working directory


	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	
	// 1. Create a directory


	err = os.Mkdir("mydir", 0755)
	if err != nil {
		fmt.Println(err)
	}

}
