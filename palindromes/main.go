/*
Determine if two strings are one insert or edit away from being equal
*/

package main

import (
	"fmt"
)

type test struct {
	str1   string
	str2   string
	result bool
}

var tests []test

func main() {

	for _, t := range tests {
		if isPalindrome(t.str1, t.str2) {
			fmt.Println("Strings" + " " + t.str1 + " " + t.str2 + " are palindromes")
		} else {
			fmt.Println("Strings" + " " + t.str1 + " " + t.str2 + " are not palindromes")
		}
	}
}

func isPalindrome(str1, str2 string) bool {
	len1 := len(str1)
	len2 := len(str2)
	diff := len1 - len2

	if diff != 0 {
		return false
	}

	for idx, ch := range str1 {
		pos := len2 - idx - 1
		if ch != int32(str2[pos]) {
			return false
		}
		if idx >= pos-1 {
			break
		}
	}
	return true
}

func init() {
	tests = make([]test, 0)
	tests = append(tests, test{"a", "a", true})
	tests = append(tests, test{"dad", "dad", true})
	tests = append(tests, test{"hannah", "hannah", true})
	tests = append(tests, test{"abe", "normal", false})
	tests = append(tests, test{"one", "ones", false})
	tests = append(tests, test{"abba", "abba", true})
	tests = append(tests, test{"", "", true})

}
