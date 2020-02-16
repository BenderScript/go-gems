// Runes, bytes and range loops gotchas

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	fmt.Printf("\n\n")
	str := "a string is an array of bytes"
	fmt.Printf("str: \"%s\"\n", str)
	fmt.Printf("str[0] type is: %T\n", str[0])
	for i, v := range str {
		fmt.Printf("But value type starting at index %d in range loop is: %T\n", i, v)
		break
	}
	fmt.Printf("\n\n")
	fmt.Printf("char type such as 'a' is: %T\n", 'a')
	fmt.Printf("Is char equal to str[0]: %v\n", 'a' == str[0])
	fmt.Printf("\n\n")
	h := "Hello, 世界"
	fmt.Printf("h: \"%s\"\n", h)
	fmt.Printf("h[7] type is: %T\n", h[7])
	fmt.Println("h[7] is only part of a rune")
	fmt.Printf("h[7] value is: %v\n", h[7])
	fmt.Printf("The complete rune is h[7:10]: %v\n", h[7:10])
	r, _ := utf8.DecodeRuneInString(h[7:])
	fmt.Println("Value of rune is:", r)
}
