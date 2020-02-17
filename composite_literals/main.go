// Composite Literals is the Go wording for explicit
// initialization of structs, slices, arrays, maps

package main

import "fmt"

type subTest struct {
	a string
	b []int
	c map[int]int
}

type test struct {
	st []subTest
}

func main() {
	// Simple Composite Literals
	aSlice := []int{1, 2, 3}
	aMap := map[int]int{1: 1, 2: 2, 3: 3}
	aArray := [...]string{"a", "b", "c"}
	fmt.Println("Slice:", aSlice)
	fmt.Println("Map: ", aMap)
	fmt.Println("Array: ", aArray)

	// A bit more complex
	st := subTest{"a string", []int{1, 2}, map[int]int{1: 1, 2: 2, 3: 3}}
	fmt.Println("Struct:", st)

	// Two-Level, complex

	t := test{[]subTest{{"a string", []int{1, 2}, map[int]int{1: 1, 2: 2, 3: 3}}, {"a string", []int{3, 4}, map[int]int{4: 4, 5: 5, 6: 6}}}}
	fmt.Println("Two-Level: ", t)
}
