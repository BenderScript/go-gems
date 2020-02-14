//Given an array of integers, return indices of the two numbers such that they add up to a specific target.
//
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
//Example:
//
//Given nums = [2, 7, 11, 15], target = 9,
//
//Because nums[0] + nums[1] = 2 + 7 = 9,
//return [0, 1].

package main

import (
	"fmt"
	"reflect"
)

type test struct {
	array   []int
	indexes []int
	sum     int
}

var tests []test

func main() {

	for _, t := range tests {
		ret := findSum(t.array, t.sum)
		if !reflect.DeepEqual(ret, t.indexes) {
			panic("Test failed")
		} else {
			fmt.Println(ret)
		}
	}

}

func findSum(arr []int, sum int) (ret []int) {
	indx := make(map[int]int, 0)
	arrMapped := false
	l := len(arr)
	indx[arr[0]] = 0
Loop:
	for i := 0; i < l-1; i++ {
		a := arr[i]
		if v, ok := indx[sum-a]; ok && (v != i) {
			ret = []int{i, v}
			break Loop
		}
		// Once array is mapped, this loop is not executed anymore
		if !arrMapped {
			for j := i + 1; j < l; j++ {
				b := arr[j]
				indx[b] = j
				if a+b == sum {
					ret = []int{i, j}
					break Loop
				}
			}
			arrMapped = true
		}
	}
	return ret
}

func init() {
	tests = append(tests, test{[]int{3, 2, 4}, []int{1, 2}, 6})
	tests = append(tests, test{[]int{10, 1, 8, 3, 4, 2}, []int{3, 4}, 7})
	tests = append(tests, test{[]int{4, 0, 9, 3, 10, 2}, []int{0, 5}, 6})
	tests = append(tests, test{[]int{1, 8, 7, 3, 5, 2}, []int{0, 4}, 6})
	tests = append(tests, test{[]int{2, 8, 0, 3, 2, 9}, []int{2, 5}, 9})
}
