package main

import (
	"fmt"
	"os"
    "bufio"
    "strings"
)

var command string

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please write second Argument after main.go")
        fmt.Println("But the Program is running ...")
	} else if len(os.Args) == 2 {
        command = os.Args[1]
        fmt.Println("The Program is running ...")
        for {
            dir, err := os.Getwd()

            if command == "pwd" {
                fmt.Println(dir)
            } else if command == "cd .." {
                err := os.Chdir("..")
                if err != nil {
                    fmt.Println(err)
                }
            } else if command == "ls" {
                files, err := os.ReadDir(".")
                if err != nil {
                    fmt.Println(err)
                }
                // fmt.Printf("%T %v", files, files) تشخیص نوع ذخیره سازی فایل ها فکر کنم اسلایس باشه :)
                for _, file := range files {
                    fmt.Println(file.Name())
                }
            }

            reader := bufio.NewReader(os.Stdin)
            command, _ = reader.ReadString('\n')
            command = strings.TrimSuffix(command, "\n")
    
            if err != nil {
                fmt.Println(err)
            }
    
            
        }
    }

    for {
        dir, err := os.Getwd()

        if command == "pwd" {
            fmt.Println(dir)
        } else if command == "cd .." {
            err := os.Chdir("..")
            if err != nil {
                fmt.Println(err)
            }
        } else if command == "ls" {
            files, err := os.ReadDir(".")
            if err != nil {
                fmt.Println(err)
            }
            // fmt.Printf("%T %v", files, files) تشخیص نوع ذخیره سازی فایل ها فکر کنم اسلایس باشه :)
            for _, file := range files {
                fmt.Println(file.Name())
            }
        }

        reader := bufio.NewReader(os.Stdin)
        command, _ = reader.ReadString('\n')
        command = strings.TrimSuffix(command, "\n")

        if err != nil {
            fmt.Println(err)
        }

        
    }

	

}







if len(os.Args) < 2 {
    fmt.Println("Please provide a command-line argument")
    return
}

arg := os.Args[1]
if arg == "" || arg == " " {
    fmt.Println("The command-line argument is empty")
    return
}

fmt.Println("The command-line argument is:", arg)

    var secondNumber *int
    if secondNumber == nil {
        fmt.Println("secondNumber is nil")
    }

    // allocate memory for secondNumber
    secondNumber = new(int)
	fmt.Println(secondNumber)
    *secondNumber = 42
    fmt.Println(*secondNumber)


// 	Package os in Golang
// 	Important Source:   https://sinalalenakhsh.notion.site/1-Chat-GPT-310ba8a7e074493e95d8c45bfe6aed86
// 	Notion Source: 		https://sinalalenakhsh.notion.site/14-Package-os-Part-1-2-3-4-2799ac556ef14fe294c8898413373dd1


package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {

	for {
		reader := bufio.NewReader(os.Stdin)

		command , _ := reader.ReadString('\n')
		command = strings.TrimSuffix(command, "\n")

		dir, err := os.Getwd()

		if err != nil {
			fmt.Println(err)
		}

		if command == "pwd" {
			fmt.Println(dir)
		} else if command == "cd .." {
			err := os.Chdir("..")
			if err != nil {
				fmt.Println(err)
			}
		} else if command == "ls" {
			files, err := os.ReadDir(".")
			if err != nil {
				fmt.Println(err)
			}
			// fmt.Printf("%T %v", files, files) تشخیص نوع ذخیره سازی فایل ها فکر کنم اسلایس باشه :)
			for _, file := range files {
				fmt.Println(file.Name())
			}
		}
	}

}
