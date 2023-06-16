// https://gobyexample.com/signals


package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

func main() {

    sigs := make(chan os.Signal, 1)

    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    done := make(chan bool, 1)

    go func() {

        sig := <-sigs
        fmt.Println()
        fmt.Println(sig)
        done <- true
    }()

    fmt.Println("awaiting signal")
    <-done
    fmt.Println("exiting")
}


/*

Sometimes we’d like our Go programs to intelligently handle Unix signals. 
For example, we might want a server to gracefully shutdown when it receives 
a SIGTERM, or a command-line tool to stop processing input if it receives 
a SIGINT. Here’s how to handle signals in Go with channels.


Go signal notification works by sending os.Signal values on a channel. 
We’ll create a channel to receive these notifications. 
Note that this channel should be buffered.


signal.Notify registers the given channel to receive notifications of the specified signals.


*/



