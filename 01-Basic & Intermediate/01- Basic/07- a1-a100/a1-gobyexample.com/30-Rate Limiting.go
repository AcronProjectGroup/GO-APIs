package main

import (
    "fmt"
    "time"
)

func main() {

    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)

    limiter := time.Tick(200 * time.Millisecond)

    for req := range requests {
        <-limiter
        fmt.Println("request", req, time.Now())
    }

    burstyLimiter := make(chan time.Time, 3)

    for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
    }

    go func() {
        for t := range time.Tick(200 * time.Millisecond) {
            burstyLimiter <- t
        }
    }()

    burstyRequests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)
    for req := range burstyRequests {
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
}



/*
Rate limiting is an important mechanism for 
controlling resource utilization and maintaining quality of service. 
Go elegantly supports rate limiting with goroutines, channels, and tickers.

محدود کردن نرخ یک مکانیسم مهم برای کنترل استفاده از منابع و حفظ کیفیت خدمات است.
گو به زیبایی از محدود کردن نرخ با گوروتین‌ها، کانال‌ها و علامت‌ها پشتیبانی می‌کند.

First we’ll look at basic rate limiting. 
Suppose we want to limit our handling of incoming requests. 
We’ll serve these requests off a channel of the same name.


This limiter channel will receive a value every 200 milliseconds. 
This is the regulator in our rate limiting scheme.



*/