// https://pkg.go.dev/strings#example-Fields

// func Fields ¶



// func Fields(s string) []string


// Fields splits the string s around each instance of one or more consecutive white space characters, as defined by unicode.IsSpace, returning a slice of substrings of s or an empty slice if s contains only white space.
// Example ¶

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
}


// Output:
// Fields are: ["foo" "bar" "baz"]

