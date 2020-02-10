/*
Binary integer order reversal. This is twist on the well-known string reversal. The idea is to swap
the actual bits.

Go has a function that provides this functionality but this is home grown

*/

package main

import (
	"fmt"
)

type test struct {
	original uint8
	expected uint8
}

var tests []test

func main() {

	for _, t := range tests {
		reversed := reverseUInt8(t)
		if reversed == t.expected {
			fmt.Printf("%d is the reverse of %d \n", t.expected, t.original)
		} else {
			panic("Not properly reversed")
		}
	}
}

func reverseUInt8(t test) (rev uint8) {
	var oneUint8, bit uint8
	oneUint8 = byte(0b00000001)

	rev = t.original
	for n := 7; n >= 4; n-- {
		lBit := (rev >> n) & oneUint8
		rBit := (rev >> (7 - n)) & oneUint8
		// Could actually be a constant
		bit = ^(oneUint8 << (7 - n)) //all ones but one bit
		rev = (rev & bit) | lBit<<(7-n)
		bit = ^(oneUint8 << n) //all ones but one bit
		rev = (rev & bit) | rBit<<n
	}
	return
}

func init() {
	tests = append(tests, test{85, 170})
	tests = append(tests, test{0, 0})
	tests = append(tests, test{255, 255})
	tests = append(tests, test{240, 15})
}
