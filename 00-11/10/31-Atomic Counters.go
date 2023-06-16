// https://gobyexample.com/atomic-counters


package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {

    var ops uint64

    var wg sync.WaitGroup

    for i := 0; i < 50; i++ {
        wg.Add(1)

        go func() {
            for c := 0; c < 1000; c++ {

                atomic.AddUint64(&ops, 1)
            }
            wg.Done()
        }()
    }

    wg.Wait()

    fmt.Println("ops:", ops)
}



/*


The primary mechanism for managing state in Go is communication over channels. 
We saw this for example with worker pools. 
There are a few other options for managing state though. 
Here we’ll look at using the sync/atomic package for atomic counters accessed by multiple goroutines.


We’ll use an unsigned integer to represent our (always-positive) counter.

A WaitGroup will help us wait for all goroutines to finish their work.


*/