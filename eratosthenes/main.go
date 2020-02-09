/*
Build a Sieve of Eratosthenes up to a limit
*/

package main

import (
	"fmt"
	"math"
)

type test struct {
	limit     float64
	composite []bool
}

var tests []test

func main() {
	for _, t := range tests {
		buildSieve(t)
		fmt.Printf("Primes (limit: %d): ", int(t.limit))
		for index, val := range t.composite {
			if !val && index > 1 {
				fmt.Printf("%d ", index)
			}
		}
		fmt.Println()
	}
}

func buildSieve(t test) {
	for i := 2; i <= int(math.Sqrt(t.limit)); i = i + 1 {
		if !t.composite[i] {
			for p := 2; ; p++ {
				next := p * i
				if next > int(t.limit) {
					break
				}
				t.composite[next] = true
			}
		}
	}
	return
}

func init() {
	tests = make([]test, 0)
	tests = append(tests, test{10, make([]bool, 11)})
	tests = append(tests, test{25, make([]bool, 26)})
	tests = append(tests, test{50, make([]bool, 51)})
	tests = append(tests, test{100, make([]bool, 101)})
}
