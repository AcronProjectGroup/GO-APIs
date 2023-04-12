/*

Bit is stand for:
    Base statement is 2
	smallest things in computer
	two statement 0/1  OR +/- for this behavier => we call it Binary دودویی


	why human use 10 base statement? look at your fingers ...

    1 bit = 2 statement = 0/1

    2 bit = 4 statement = 00  1  0
                          01  2
                          10  3
                          11  4

    3 bit = 8 statement = 000  1
                          001  2
                          010  3
                          011  4
                          100  5
                          101  6
                          110  7
                          111  8

    4 bit = 16  statement 0000
    5 bit = 32  statement 00000
    6 bit = 64  statement 000000
    7 bit = 128 statement 0000000
    8 bit = 255 statement 00000000 = 0 ->   11111111 -> 254 = 255 -> 1 Byte




BYTE is

    Per 1 Byte is 8 Bit (Lined up)

               Hardware
               --------
               ||||||||
        byte [ 00000000 ] = 8 bit
               87654321
               ||||||||
               ||||||||_ 2^0 = 1
               |||||||
               |||||||_ 2^1 = 2
               ||||||
               ||||||_ 2^2 = 4
               |||||
               |||||_ 2^3 = 8
               ||||
               ||||_ 2^4 = 16
               |||
               |||_ 2^5 = 32
               ||
               ||_ 2^6 = 64
               |
               |_ 2^7 = 128

               128 + 64 + 32 + 16 + 8 + 4 + 2 + 1 = 255 statement capacity


Convertion:
    https://convertlive.com/u/convert/megabytes/to/bytes

*/

// package main

// import "fmt"

// func main() {
// 	var Harchizi byte = 255
// 	fmt.Printf("%08b \n", Harchizi)
// }

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		operator1()
	} else if len(os.Args) > 2 && os.Args[1] == "converttobinary" {

        //------------------------------------------------------------------------ 
        args := os.Args[2:]                                                        /* main.go converttobinary fromhere to more .... */

        //------------------------------------------------------------------------ /* this is add spaces between words to slice of bytes */
		bytes := []byte(strings.Join(args, " "))                                   /* fmt.Println(bytes) */
                                                                                             

        //------------------------------------------------------------------------
		// var result string                                                          /* for attach every index of slice */
		// for _, value := range bytes {
		// 	result = result + strconv.Itoa(int(value))
		// }
		// fmt.Printf("Result of conversion to Bytes: %v\n", result)

        //------------------------------------------------------------------------
        var convertToBinary string 
        for _, value := range bytes {
            convertToBinary = convertToBinary + strconv.FormatInt(int64(value), 2) /* Convert to binary and append to convertToBinary */
        }
		fmt.Printf("Result of conversion to Binaries: %v\n", convertToBinary)



	}

	operator1()

}

func operator1() {

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
		} else if command == "finish" || command == "exit" {
			return
		} else if command == "converttobinary" {
			for _, arg := range os.Args[2:] {
				num, err := strconv.Atoi(arg)
				if err != nil {
					fmt.Printf("error converting %s to integer: %v\n", arg, err)
					continue
				}
				fmt.Printf("%b\n", num)
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
