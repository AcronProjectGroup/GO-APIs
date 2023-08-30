package main

import (
    "fmt"
    "time"
)

func main() {

    ticker := time.NewTicker(500 * time.Millisecond)
    done := make(chan bool)

    go func() {
        for {
            select {
            case <-done:
                return
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
            }
        }
    }()

    time.Sleep(1600 * time.Millisecond)
    ticker.Stop()
    done <- true
    fmt.Println("Ticker stopped")
}

/*

Timers are for when you want to do something once in the future 
- tickers are for when you want to do something repeatedly at regular intervals. 
Here’s an example of a ticker that ticks periodically until we stop it.

تایمرها برای زمانی هستند که می خواهید کاری را یک بار در آینده انجام دهید - 
تیک تیک ها برای زمانی هستند که می خواهید کاری را به طور مکرر در فواصل زمانی منظم انجام دهید. 
در اینجا نمونه ای از تیک تیک است که به صورت دوره ای تیک می زند تا زمانی که آن را متوقف کنیم.

*/