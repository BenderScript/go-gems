//
// https://golang.org/pkg/reflect/#StringHeader
// https://tylerchr.blog/golang-arbitrary-memory/
// https://stackoverflow.com/questions/43766471/accessing-a-memory-address-from-a-string-in-go
// https://stackoverflow.com/questions/47859348/convert-uintptr-from-reflection-into-pointer
// https://stackoverflow.com/questions/51187973/how-to-create-an-array-or-a-slice-from-an-array-unsafe-pointer-in-golang
// https://golang.org/pkg/unsafe/

package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	oStr := "all is good"
	fmt.Printf("String is: %q \n", oStr)

	// Let's get a untyped pointer to the underlining string representation
	p := unsafe.Pointer(&oStr)
	// https://golang.org/pkg/reflect/#StringHeader
	sh := (*reflect.StringHeader)(p)
	fmt.Printf("Address of internal string representation: %p \n", sh)
	fmt.Printf("Pointer to actual string is: %p \n", unsafe.Pointer(sh.Data))
	end := sh.Data + uintptr(sh.Len)
	// Walk the memory addresses and print each char individually
	fmt.Print("Walking memory and printing string...")
	for i := sh.Data; i < end; i = i + uintptr(1) {
		pb := (*uint8)(unsafe.Pointer(i))
		fmt.Printf("%s", string(*pb))
	}
}
