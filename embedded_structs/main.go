package main

import "fmt"

type topLevel struct {
	level2
}

type level2 struct {
	string2     string
	printString func(string) bool
	level3
}

type level3 struct {
	string3 string
}

func main() {
	t := new(topLevel)
	t.printString = func(a string) bool {
		fmt.Println(a)
		return true
	}
	t.string3 = "level3"
}
