// https://gobyexample.com/file-paths


package main

import (
    "fmt"
    "path/filepath"
    "strings"
)

func main() {

    p := filepath.Join("dir1", "dir2", "filename")
    fmt.Println("p:", p)

    fmt.Println(filepath.Join("dir1//", "filename"))
    fmt.Println(filepath.Join("dir1/../dir1", "filename"))

    fmt.Println("Dir(p):", filepath.Dir(p))
    fmt.Println("Base(p):", filepath.Base(p))

    fmt.Println(filepath.IsAbs("dir/file"))
    fmt.Println(filepath.IsAbs("/dir/file"))

    filename := "config.json"

    ext := filepath.Ext(filename)
    fmt.Println(ext)

    fmt.Println(strings.TrimSuffix(filename, ext))

    rel, err := filepath.Rel("a/b", "a/b/t/file")
    if err != nil {
        panic(err)
    }
    fmt.Println(rel)

    rel, err = filepath.Rel("a/b", "a/c/t/file")
    if err != nil {
        panic(err)
    }
    fmt.Println(rel)
}


/*

$ go run file-paths.go
p: dir1/dir2/filename
dir1/filename
dir1/filename
Dir(p): dir1/dir2
Base(p): filename
false
true
.json
config
t/file
../c/t/file

*/

// ---------------------------

/*

The filepath package provides functions to parse and construct file paths 
in a way that is portable between operating systems; 

dir/file on Linux vs. dir\file on Windows, for example.

*/