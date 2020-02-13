/*
Caesar cipher for ASCII chars with an alphabet of 26 chars
*/

package main

import (
	"fmt"
)

type test struct {
	str1  string
	str2  string
	shift int32
}

var tests []test

func main() {

	for _, t := range tests {
		if caesarCipher(t.str1, t.str2, t.shift) {
			fmt.Println(t.str2 + " is proper cipher of " + t.str1)
		} else {
			fmt.Println(t.str2 + " is not proper cipher of " + t.str1)
		}
	}
}

// We assume only lowercase ASCII and an alphabet of 26 letters
func caesarCipher(str1, str2 string, shift int32) bool {

	var result []rune
	for _, elem := range str1 {
		if elem == ' ' {
			result = append(result, ' ')
		} else {
			shift = shift % 26
			newElem := elem + shift
			if newElem > 'z' {
				newElem = newElem - 26
			}
			result = append(result, newElem)
		}
	}
	cipher := string(result)
	if cipher == str2 {
		return true
	}
	return false
}

func init() {
	tests = make([]test, 0)
	tests = append(tests, test{"interview question", "nsyjwanjb vzjxynts", 5})
	tests = append(tests, test{"caesar cipher", "igkygx iovnkx", 6})
	tests = append(tests, test{"modulus is elegant", "tvkbsbz pz lslnhua", 7})
	tests = append(tests, test{"caesar cipher", "caesar cipher", 26})
	tests = append(tests, test{"caesar cipher", "geiwev gmtliv", 30})
}
