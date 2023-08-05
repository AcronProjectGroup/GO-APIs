/*


What are they?
	Goroutines are lightweight threads created and managed by the Go runtime.
	Channels are pipes that carry values of a specific type.

Why are they useful?
	Goroutines allow functions to be executed concurrently, without needing to deal
	with the complications of operating system threads. Channels allow goroutines to
	produce results asynchronously.

How are they used?
	Goroutines are created using the go keyword. Channels are defined as data types.

Are there any pitfalls or limitations?
	Care must be taken to manage the direction of channels. Goroutines that share data
	require additional features, which are described in Chapter 14.

Are there any alternatives?
	Goroutines and channels are the built-in Go concurrency features, but some
	applications can rely on a single thread of execution, which is created by default to
	execute the main function.



Problem:

Execute a function asynchronously:
	Create a goroutine


Produce a result from a function executed asynchronously:
	Use a channel


Send and receive values using a channel:
	Use arrow expressions

Indicate that no further values will be sent over a channel:
	Use the close function

Enumerate the values received from a channel:
	Use a for loop with the range keyword

Send or receive values using multiple channels:
	Use a select statement







*/


package main