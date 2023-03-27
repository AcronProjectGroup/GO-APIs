// Get User Input Part2

// Notion Source : https://sinalalenakhsh.notion.site/9-every-ways-get-user-input-in-Golang-ac2c758f26b04902b0c706194dc77c68

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter your age: ")

    age, _ := reader.ReadString('\n')

	age = strings.TrimSuffix(age, "\n")
	
    fmt.Printf("You are %s years old.\n", age)


}

/*

	1. Using the bufio package:

	```
	package main

	import (
		"bufio"
		"fmt"
		"os"
	)

	func main() {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your age: ")
		age, _ := reader.ReadString('\n')
		fmt.Printf("You are %s years old.\n", age)
	}
	```

	why in this code " years old" go to new line ?

	package main

	import (
		"bufio"
		"fmt"
		"os"
	)

	func main() {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter your age: ")

		age, whatIsIt := reader.ReadString('\n')

		fmt.Printf("You are %s years old.", age)

		fmt.Println(whatIsIt)
	}

	In the provided code, the string " years old." does not necessarily go to a new line. It depends on the input entered by the user.

	The `reader.ReadString('\n')` function reads the input from the user until it encounters a newline character ('\n') and returns the input string including the newline character.

	So if the user enters their age and presses "Enter", the input string returned by `reader.ReadString('\n')` will contain the age followed by a newline character. When the input string is printed using `fmt.Printf("You are %s years old.", age)`, the newline character in the input string causes the output to start on a new line.

	If you want to avoid this behavior, you can trim the newline character from the input string using the `strings.TrimSuffix()` function like this:

	```
	age = strings.TrimSuffix(age, "\n")
	fmt.Printf("You are %s years old.", age)

	```

	This will remove the newline character from the input string and the output will not start on a new line.
*/