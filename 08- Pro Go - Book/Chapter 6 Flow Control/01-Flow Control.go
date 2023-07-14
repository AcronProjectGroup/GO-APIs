/*

Question			Answer
--------			----------------------------------------
What is it?			Flow control allows a programmer to selectively execute statements.

Why is it useful?	Without flow control, an application executes a series of code statements in
					sequence and then exits. Flow control allows this sequence to be altered, deferring
					the execution of some statements and repeating the execution of others.

How is it used?		Go supports flow control keywords, including if, for, and switch, each of which
					controls the flow of execution differently.

Are there any		Go introduces unusual features for each of its flow control keywords that offer
pitfalls or			additional features, which must be used with care.
limitations?

Are there any		No. Flow control is a fundamental language feature.
alternatives?

*/

package main

import "fmt"

func main() {
	kayakPrice := 275.00
	fmt.Println("Price:", kayakPrice)
}



