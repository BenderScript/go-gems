/*
URLfy a string
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
		if urlFy(t.str1, t.str2) {
			fmt.Println("URLfy successful")
		} else {
			fmt.Println("URLfy not successful")
		}
	}
}

func urlFy(str1, str2 string) bool {

	spc := "%20"
	l := len(str1)
	for i := 0; ; i++ {
		if i == l {
			break
		}
		if str1[i] == ' ' {
			str1 = str1[:i] + spc + str1[i+1:]
			i, l = i+2, l+2
		}
	}

	return str1 == str2
}

func init() {
	tests = make([]test, 0)
	tests = append(tests, test{"a new url", "a%20new%20url", false})
	tests = append(tests, test{"space in the end ", "space%20in%20the%20end%20", true})
	tests = append(tests, test{"two  spaces", "two%20%20spaces", false})

}
