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
	fmt.Printf("b[len(b):] returns empty string: %q \n", b[len(b):])
	fmt.Println("======")
	c := make([]int, 3)
	fmt.Printf("c is : %T \n", c)
	fmt.Println("c[len(c):len(c)] returns empty slice: ", c[len(c):len(c)])
	fmt.Println("======")
	d := make([]int, 3, 5)
	fmt.Printf("d is %T with len (%d) and cap (%d) \n", d, len(d), cap(d))
	fmt.Println("d[5:5] is not out of bounds: ", d[5:5])
	fmt.Println("because d is a slice of int and cap is: ", cap(d))
	fmt.Printf("But d[5:] is out of bounds since d missing high index defaults to length : %d \n", len(d))
	_ = d[5:]
}
