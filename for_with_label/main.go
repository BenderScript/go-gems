package main

import "fmt"

func main() {

Next:
	// No code between label and for loop
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if j == 1 {
				fmt.Println("Continue to next outer iteration")
				continue Next
			}
			if i == 1 {
				fmt.Println("Breaking out of all loops")
				break Next
			}
		}
	}
	fmt.Println("First Statement after Next Label")
}
