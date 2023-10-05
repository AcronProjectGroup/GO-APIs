package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	fmt.Println("Do you want to open the browser? (yes/no)")

	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)

	if strings.ToLower(userInput) == "yes" {
		openBrowser()
	} else {
		fmt.Println("Okay, not opening the browser.")
	}
}

func openBrowser() {
	var command string

	switch runtime.GOOS {
	case "darwin":
		command = "open"
	case "windows":
		command = "start"
	case "linux":
		command = "xdg-open"
	default:
		fmt.Println("Unsupported platform.")
		return
	}

	exec.Command(command, "http://example.com").Start()
}

