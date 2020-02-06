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
		if isAway(t.str1, t.str2) {
			fmt.Println("Strings" + " " + t.str1 + " " + t.str2 + " are one away")
		} else {
			fmt.Println("Strings" + " " + t.str1 + " " + t.str2 + " are not one away")
		}
	}
}

// Reuse most of the isAway function
func isAway(str1, str2 string) bool {
	var charMap map[rune]int
	charMap = make(map[rune]int)

	len1 := len(str1)
	len2 := len(str2)
	diff := len1 - len2

	if diff > 1 || diff < -1 {
		return false
	}

	for _, char := range str1 {
		charMap[char] += 1
	}
	for _, char := range str2 {
		charMap[char] -= 1
		if charMap[char] == 0 {
			delete(charMap, char)
		}
	}
	if len(charMap) >= 0 && len(charMap) <= 1 {
		return true
	}
	return false
}

func init() {
	tests = make([]test, 0)
	tests = append(tests, test{"pepe", "pep", false})
	tests = append(tests, test{"hannah", "hannah", true})
	tests = append(tests, test{"abe", "normal", false})
	tests = append(tests, test{"one", "oness", true})
	tests = append(tests, test{"a", "", true})
}
