package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Element struct {
	Data byte
	Prev *Element
	Next *Element
}

type Command string

const (
	INSERT    Command = "INSERT"
	LEFT      Command = "LEFT"
	RIGHT     Command = "RIGHT"
	BACKSPACE Command = "BACKSPACE"
)

func main() {

	var lineCount int
	fmt.Scan(&lineCount)

	reader := bufio.NewReader(os.Stdin)
	for index := 0; index < lineCount; index++ {

		line, _ := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		
		fields := strings.Fields(line)
		switch Command(fields[0]) {
		case INSERT:
			fmt.Println("INSERT")
		case LEFT:
			fmt.Println("LEFT")
		case RIGHT:
			fmt.Println("RIGHT")
		case BACKSPACE:
			fmt.Println("BACKSPACE")
		default:
			panic("Invalid Command Given!")
		}
	}

}

/* Output:

go run main.go < test.txt 
INSERT
INSERT
INSERT
LEFT
LEFT
RIGHT
BACKSPACE
*/