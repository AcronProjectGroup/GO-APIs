package main

import "fmt"

func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")

    <-done
}
/*

Closing a channel indicates that no more values will be sent on it. 
This can be useful to communicate completion to the channel’s receivers.

In this example we’ll use a jobs channel to communicate work to be done 
from the main() goroutine to a worker goroutine. When we have no more jobs 
for the worker we’ll close the jobs channel.



*/