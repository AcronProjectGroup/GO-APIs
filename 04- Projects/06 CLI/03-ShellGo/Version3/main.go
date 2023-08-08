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

func (E *Element) GetFirstElement() *Element {
	Current := E
	for {
		if Current.Prev == nil {
			break
		}
		Current = Current.Prev
	}
	return Current
}

func (E *Element) PrintElement() {
	Current := E.GetFirstElement()
	
	var StringBuilder strings.Builder

	for {
		StringBuilder.WriteByte(Current.Data)
		if Current.Next == nil {
			break
		} else {
			Current = Current.Next
		}
	}

	fmt.Println(StringBuilder.String())
}

func main() {

	var lineCount int
	fmt.Scan(&lineCount)

	reader := bufio.NewReader(os.Stdin)
	curser := &Element{Data: '|'}

	for index := 0; index < lineCount; index++ {

		line, _ := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		
		fields := strings.Fields(line)
		switch Command(fields[0]) {
		case INSERT:
			// fmt.Println("INSERT")

			element1 := &Element{Data: []byte(fields[1])[0]}

			if curser.Prev != nil {
				element1.Prev = curser.Prev
				curser.Prev.Next = element1
			}

			curser.Prev = element1
			element1.Next = curser

		case LEFT:
			// fmt.Println("LEFT")
			if curser.Prev != nil {
				curser.Prev.Data, curser.Data = curser.Data , curser.Prev.Data
				curser = curser.Prev
			}

		case RIGHT:
			// fmt.Println("RIGHT")
			if curser.Next != nil {
				curser.Next.Data, curser.Data = curser.Data , curser.Next.Data
				curser = curser.Next
			}
		case BACKSPACE:
			// fmt.Println("BACKSPACE")
			if curser.Prev != nil {
				if curser.Prev.Prev != nil {
					curser.Prev.Prev.Next = curser
					curser.Prev = curser.Prev.Prev
				} else {
					curser.Prev = nil
				}
			}
		default:
			panic("Invalid Command Given!")
		}
	}


	
	curser.PrintElement()
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