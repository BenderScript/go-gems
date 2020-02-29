// A hidden gem that causes bugs ans is used in many interviews
// https://golang.org/ref/spec#Slice_expressions

package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	b := "example"
	fmt.Printf("b is: %q \n", b)
	fmt.Printf("b[len(b):] is not out of range: %q \n", b[len(b):])
	c := make([]int, 3)
	fmt.Println("c[len(c):len(c)] is not out of bounds: ", c[len(c):len(c)])
	d := make([]int, 3, 5)
	fmt.Println("d is: ", d)
	fmt.Println("d[5:5] is not out of bounds: ", d[5:5])
	fmt.Println("Because it is d slice of int and cap is: ", cap(d))
	fmt.Println("But d[5:] is out of bounds since d missing high index defaults to the length ")
	_ = d[5:]
}
