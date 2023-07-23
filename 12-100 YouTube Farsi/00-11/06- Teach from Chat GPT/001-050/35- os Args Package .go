
// os Args Package 

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var index2 float64

	index2, _ = strconv.ParseFloat(os.Args[1], 64)
				// strconv.Atoi() for Integer
	fmt.Printf("%v -->>  %T \n", index2, index2)


}























// /*
// 	Package os in Golang
// 	Important Source:   https://sinalalenakhsh.notion.site/1-Chat-GPT-310ba8a7e074493e95d8c45bfe6aed86
// 	Notion Source: 		https://sinalalenakhsh.notion.site/14-Package-os-Part-1-2-3-4-2799ac556ef14fe294c8898413373dd1

// */

// package main

// import (
// 	"fmt"
// 	"os"
// 	"bufio"
// 	"strings"
// )

// func main() {

// 	for {
// 		reader := bufio.NewReader(os.Stdin)

// 		fmt.Println(" *** in Go Program *** ")

// 		command , _ := reader.ReadString('\n')
// 		command = strings.TrimSuffix(command, "\n")

// 		dir, err := os.Getwd()

// 		if err != nil {
// 			fmt.Println(err)
// 		}

// 		if command == "pwd" {
// 			fmt.Println(dir)
// 		} else if command == "cd .." {
// 			err := os.Chdir("..")
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 		} else if command == "ls" {
// 			files, err := os.ReadDir(".")
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 			// fmt.Printf("%T %v", files, files) تشخیص نوع ذخیره سازی فایل ها فکر کنم اسلایس باشه :)
// 			for _, file := range files {
// 				fmt.Println(file.Name())
// 			}
// 		}
// 	}

// }



package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("It's over", os.Args[1])
}