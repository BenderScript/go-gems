// Go lang Closures.
// Creates a closure function and passes it to another function as a parameter.
// This function in turn calls the closure.

package main

import (
	"fmt"
)

func main() {

	closureA := createClosure(1)
	useClosure(2, closureA)
	useClosure(5, closureA)
}

// createClosure creates two closures. One around a variable defined within the function
// and another around a function parameter
func createClosure(closureInput int) func(int) int {
	fmt.Println("Create Closure")
	closureVar := -10
	fmt.Println(fmt.Sprintf("%-20v : %d", "closureInput ", closureInput))
	fmt.Println(fmt.Sprintf("%-20v : %d", "closureVar ", closureVar))
	return func(a int) int {
		closureInput += a
		closureVar -= a
		fmt.Println(fmt.Sprintf("%-20v : %d", "closureInput ", closureInput))
		fmt.Println(fmt.Sprintf("%-20v : %d", "closureVar ", closureVar))
		return closureInput + closureVar
	}
}

// useClosure takes a function with a closure var and use it.
// It is a dramatic way of showing closures since this function has absolutely no
// reference to the "closed" variable
func useClosure(a int, f func(int) int) int {
	fmt.Println("Use Closure, input: ", a)
	return f(a)
}
