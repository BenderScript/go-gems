/*
Determine is a string is a rotation of another
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
		if isRotation(t.str1, t.str2) {
			fmt.Println("Strings" + " " + t.str1 + " " + t.str2 + " are rotations")
		} else {
			fmt.Println("Strings" + " " + t.str1 + " " + t.str2 + " are not rotations")
		}
	}
}

func isRotation(str1, str2 string) bool {

	len1 := len(str1)
	len2 := len(str2)
	if len1 != len2 {
		return false
	}
	j := 0
	match := false
	for i := 0; i < len1; i++ {
		if str1[i] == str2[j] {
			match = true
		}
		if match {
			for x, y := i+1, j+1; y < len2; x, y = x+1, y+1 {
				if str1[x] != str2[y] {
					match = false
					break
				}
				if x == len1-1 {
					x = -1
				}
			}
			if match {
				break
			}
		}
	}
	return match
}

func init() {
	tests = make([]test, 0)
	tests = append(tests, test{"pepe", "epep", true})
	tests = append(tests, test{"baby", "byba", true})
	tests = append(tests, test{"pepe", "ppee", false})
	tests = append(tests, test{"a", "a", true})
	tests = append(tests, test{"pepe", "pepe", true})
}
