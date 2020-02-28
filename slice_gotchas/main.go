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
	a := make([]int, 3, 5)
	fmt.Println("a is: ", a)
	fmt.Println("a[5:5] is not out of bounds: ", a[5:5])
	fmt.Println("Because it is a slice of int and cap is: ", cap(a))
	fmt.Println("But a[5:] is out of bounds since a missing high index defaults to the length ")
	_ = a[5:]
}
