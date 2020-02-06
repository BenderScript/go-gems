/*
Determine is a string only has unique characters
*/

package main

import "fmt"

var strings []string
var charMap map[rune]bool

func main() {
	unique := true
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	for _, str := range strings {
		for _, char := range str {
			if _, ok := charMap[char]; ok {
				fmt.Println("string " + str + " does not have unique chars")
				unique = false
				break
			} else {
				charMap[char] = true
			}
		}
		if unique {
			fmt.Println("string " + str + " has unique chars")
		}
		unique = true
	}
}

func init() {
	strings = append(strings, []string{"unique", "chars", "pepe", ""}...)
	charMap = make(map[rune]bool)
}
