/*

Problem								SolutionListing
----------------------------		---------------------------------
Store a fixed number of values		Use an array

Compare arrays						Use the comparison operators

Enumerate an array					Use a for loop with the range keyword

Store a variable number of values	Use a slice

Append an item to a slice			Use the append function

Create a slice from an existing		Use a range
array or select elements from a
slice

Copy elements into a slice			Use the copy function

Delete elements from a slice		Use the append function with ranges that omit the
									elements to remove

Enumerate a slice					Use a for loop with the range keyword31
Sort the elements in a slice		Use the sort package32
Compare slices						Use the reflect package33, 34

Obtain a pointer to the array		Perform an explicit conversion to an array type
underlying a slice					whose length is less than or equal to the number of					
									elements in the slice

Store key-value pairs				Use a map

Remove a key-value pair from a		Use the delete function
map
Enumerate the contents of a map		Use a for loop with the range keyword42, 43

Read byte values or characters		Use the string as an array or perform an explicit
from a string						conversion to the []rune type


Enumerate the characters in a		Use a for loop with the range keyword
string

Enumerate the bytes in a string		Perform an explicit conversion to the []byte type
									and use a for loop with the range keyword

*/
package main
