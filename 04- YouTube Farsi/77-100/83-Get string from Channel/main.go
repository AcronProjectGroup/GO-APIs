package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// "time"
)

func GetUserInput(Channel chan string)  {
	userInput := bufio.NewReader(os.Stdin)
	finalUserInput2, _ := userInput.ReadString('\n')
	finalUserInput2 = strings.TrimSuffix(finalUserInput2, "\n")
	finalUserInput := finalUserInput2
	Channel <- finalUserInput
}

func PrintInAnotherTimeLine(Channel <- chan string) {
	fromChannel := <-Channel
	fmt.Println(fromChannel)
}

func main()  {

	aChannel := make(chan string, 1)

	go GetUserInput(aChannel)

	PrintInAnotherTimeLine(aChannel)


	// time.Sleep(time.Second * 5)
	


}