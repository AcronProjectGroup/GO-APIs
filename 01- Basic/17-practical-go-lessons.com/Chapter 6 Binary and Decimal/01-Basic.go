/*
Technical concepts covered

	# Binary
	# Bits, Bytes
	# Decimal
	# Encoding, Encoder
	# RAM

Numerals, numbers, and quantity:

	# Numeral : a symbol or a group of characters that represent a number1
	# A number represent a quantity, something that we can count.
	# A numeral is a representation of a number.
	# A numeration system is a set of rules that allow us to count things and to represent quantities.

“123” represent this quantity in the decimal system

	The numeral is “123”

“1111011” represent this quantity in the binary system

	The numeral is “1111011”

“7B” represent this quantity in the hexadecimal system

	The numeral is “7B”

“173” represent this quantity in the octal system

	The numeral is “173”

Etymology and symbols:

	Decimal comes from the Latin word “Decimus” which means “tenth”.
	Whereas binary comes from the Latin “bini” which means “two together”.

	The etymology of those two words give us a hint about how those systems are constructed :
	    # The binary system uses two symbols which are 0 and 1
	    # The decimal system uses ten symbols which are 0, 1, 2, 3, 4, 5, 6, 7, 8, 9.

The storage capacity of 8 bits (one byte)

	11111111binary​=( (1×2^7) + (1×2^6) + (1×2^5) + (1×2^4) + (1×2^3) + (1×2^2) + (1×2^1) + (1×2^0) )decimal​
		          =255decimal
	With 8 bits, we can store all the numbers between 0decimal0decimal​ and 255decimal255decimal​.
	8 bits is called a byte.
	From 0 decimal​ to 255 decimal​ there are 256 numbers.

How to store images, videos,... ?

	The answer is simple:
	at the end, even photos and movies will be stored using zeros and ones!
	We will convert them into binary.
	This job will be performed by specialized programs called encoders.
	An encoder will take as input a file from a specific format and will convert it into a destination format.
	In our case, the destination is binary.


	We do not need to write those programs; they are all provided by Go.
	You need to understand that every file or chunk of data is stored under the hood using binary digits (bits).
	The binary representation is just hidden from us.

32 vs. 64 bits systems

	#1 Programs need memory:
			Most programs need to store and access memory. If you write a program that will display “42” on the screen. Behind the scene, the number 42 needs to be stored somewhere. The system will also need to fetch it from memory.
			Remember that there are two types of memory :
				The central memory: ROM and RAM
				The auxiliary memory: Hard drive, USB keys...
			Here we consider only the central memory.

	#2 A memory address is a number:
			For a processor an address is a number.
			A memory address is like a postal address.
			It identifies precisely a point in memory space.
			Processors will store addresses on registers.
			A register is a place inside the processor where an address can be saved for later use.

	#3 Memory addressable: a limited resource
			Registers have a limited capacity;
			they can store addresses of a specific size (written in bits).
			For a 16 bit processor, the maximum capacity of those registers is... 16 bits.
			For 32 bits processors, the maximum capacity is 32 bits.
			The same reasoning applies to 64 bits.
			The maximum register capacity will define how much memory we can address.


The relation between the number of addresses possible and RAM size:
	RAM is a hardware component composed of memory cells.
	Bits are stored in cells.
	In general, RAM is said to be byte-addressable.
	It means that the system can fetch data 8 bits at a time.
	We have seen that the size of a memory address is limited by the size of registers.
	A 32-bit system can only handle addresses that are composed of 32 bits.
	Each bit is composed either of a 0 or a 1, which makes 232≈4232≈4 billion possibilities.

With 4 billion addresses, how many bytes can we address?
	=> 4 billion bytes !

We can convert this quantity into Gigabytes.
	1 Gigabyte = 1 billion bytes
	So, 4 billion bytes = 4 Gigabytes





	
	


















*/
package main

import (
	"fmt"
	"math"
)

func main()  {
	myNumber := math.Pow(2, 32)
	fmt.Printf("%.f\n", myNumber)
	myNumber = math.Pow(2, 64)
	fmt.Printf("%.f\n", myNumber)
}