/*
Determine is a string is a permutation of another
*/

package main

import "reflect"

type test struct {
	slc1 []string
	slc2 []string
	exp  []string
}

var tests []test

func main() {

	if !reflect.DeepEqual(copySlice(tests[0].slc1), tests[0].exp) {
		panic(tests[0].slc1)
	}
	if !reflect.DeepEqual(cutSlice(tests[1].slc1), tests[1].exp) {
		panic(tests[1].slc1)
	}
	if !reflect.DeepEqual(expandSlice(tests[2].slc1, 2), tests[2].exp) {
		panic(tests[2].slc1)
	}
	if !reflect.DeepEqual(extendSlice(tests[3].slc1, 2), tests[3].exp) {
		panic(tests[3].slc1)
	}
}

func extendSlice(source []string, positions int) []string {
	source = append(source, make([]string, positions)...)
	return source
}

func expandSlice(source []string, positions int) []string {
	// a = append(a[:i], append(make([]T, j), a[i:]...)...)
	l := len(source)
	source = append(source[:l/2], append(make([]string, positions), source[l/2:]...)...)
	return source
}

func cutSlice(source []string) []string {
	l := len(source)
	b := append(source[:l/2], source[l/2+1:]...)
	return b
}

func copySlice(source []string) []string {
	b := make([]string, len(source))
	copy(b, source)
	return b
}

func init() {
	tests = make([]test, 0)
	tests = append(tests, test{[]string{"a", "b", "c", "d"}, []string{"d", "e", "f", "g"}, []string{"a", "b", "c", "d"}})
	tests = append(tests, test{[]string{"a", "b", "c", "d"}, []string{"d", "e", "f", "g"}, []string{"a", "b", "d"}})
	tests = append(tests, test{[]string{"a", "b", "c", "d"}, []string{"d", "e", "f", "g"}, []string{"a", "b", "", "", "c", "d"}})
	tests = append(tests, test{[]string{"a", "b", "c", "d"}, []string{"d", "e", "f", "g"}, []string{"a", "b", "c", "d", "", ""}})
}
