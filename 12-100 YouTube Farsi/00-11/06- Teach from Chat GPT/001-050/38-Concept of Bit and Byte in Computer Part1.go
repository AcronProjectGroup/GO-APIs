/*

Bit is stand for:
    Base statement is 2
	smallest things in computer
	two statement 0/1  OR +/- for this behavier => we call it Binary دودویی


	why human use 10 base statement? look at your fingers ...

    1 bit = 2 statement = 0/1

    2 bit = 4 statement = 00  1  0
                          01  2
                          10  3
                          11  4

    3 bit = 8 statement = 000  1
                          001  2
                          010  3
                          011  4
                          100  5
                          101  6
                          110  7
                          111  8

    4 bit = 16  statement 0000
    5 bit = 32  statement 00000
    6 bit = 64  statement 000000
    7 bit = 128 statement 0000000
    8 bit = 255 statement 00000000 = 0 ->   11111111 -> 254 = 255 -> 1 Byte




BYTE is

    Per 1 Byte is 8 Bit (Lined up)

               Hardware
               --------
               ||||||||
        byte [ 00000000 ] = 8 bit
               87654321
               ||||||||
               ||||||||_ 2^0 = 1
               |||||||
               |||||||_ 2^1 = 2
               ||||||
               ||||||_ 2^2 = 4
               |||||
               |||||_ 2^3 = 8
               ||||
               ||||_ 2^4 = 16
               |||
               |||_ 2^5 = 32
               ||
               ||_ 2^6 = 64
               |
               |_ 2^7 = 128

               128 + 64 + 32 + 16 + 8 + 4 + 2 + 1 = 255 statement capacity


Convertion:
    https://convertlive.com/u/convert/megabytes/to/bytes

*/


package main

import "fmt"

func main() {
	var Harchizi byte = 255
	fmt.Printf("%08b \n", Harchizi)
}
