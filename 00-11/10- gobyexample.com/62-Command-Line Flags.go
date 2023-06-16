// https://gobyexample.com/command-line-flags


package main

import (
    "flag"
    "fmt"
)

func main() {

    wordPtr := flag.String("word", "foo", "a string")

    numbPtr := flag.Int("numb", 42, "an int")
    forkPtr := flag.Bool("fork", false, "a bool")

    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    flag.Parse()

    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *forkPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
}

/*

Command-line flags are a common way to specify options 
for command-line programs. For example, in wc -l the -l is a command-line flag.


Command-line flags are a common way to specify options for command-line programs. 

For example, in wc -l the -l is a command-line flag.

Basic flag declarations are available for string, integer, and boolean options. 
Here we declare a string flag word with a default value "foo" and a short description. 
This flag.String function returns a string pointer (not a string value); 
we’ll see how to use this pointer below.

*/