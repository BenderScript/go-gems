package main

import "fmt"

// Unpack operator breaks a slice into individual components. it ends with ... like slice...
// Pack operator combine individual elements like integers into a slice. It starts with ... like ...Type
// Effective Go talks about append but wish it would spend just a bit more time on the subject.

var exNum int

func main() {

	variadicA([]string{"a", "b", "c"}...)
	variadicA("d", "e", "f")

	// The mythical variadic interface{}
	variadicInterface(1.0, "a", 2, []int{1, 2})

}

func variadicA(elems ...string) {
	exNum++
	fmt.Printf("Example number: %d", exNum)
	str := ""
	for _, e := range elems {
		str += e
	}
	fmt.Println("Output: ", str)
}

func variadicInterface(elems ...interface{}) {
	exNum++
	fmt.Printf("Example number: %d", exNum)
	for _, e := range elems {
		switch e.(type) {
		case float64:
			fmt.Println("Type is float64")
		case int:
			fmt.Println("Type is Integer")
		case string:
			fmt.Println("Type is String")
		case []int:
			fmt.Println("Type is []int")
		default:
			fmt.Printf("Catch-all. Type is %T", e)
		}
	}
}
