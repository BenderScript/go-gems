/*
Determine is a string is a permutation of another
*/

package main

import "fmt"

type test struct {
	str1   string
	str2   string
	result bool
}

var tests []test

func main() {

	for _, t := range tests {
		if isPermutation(t.str1, t.str2) {
			fmt.Println("Strings" + " " + t.str1 + " " + t.str2 + " are permutations")
		} else {
			fmt.Println("Strings" + " " + t.str1 + " " + t.str2 + " are not permutations")
		}
	}
}

func isPermutation(str1, str2 string) bool {
	var charMap map[rune]int
	charMap = make(map[rune]int)

	for _, char := range str1 {
		charMap[char] += 1
	}
	for _, char := range str2 {
		charMap[char] -= 1
		if charMap[char] == 0 {
			delete(charMap, char)
		}
	}
	if len(charMap) == 0 {
		return true
	}
	return false
}

func init() {
	tests = make([]test, 0)
	tests = append(tests, test{"pepe", "pep", false})
	tests = append(tests, test{"hannah", "hannah", true})
	tests = append(tests, test{"abe", "normal", false})
}
