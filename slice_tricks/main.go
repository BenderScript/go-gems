/*
Slice tricks.

Implementations of https://github.com/golang/go/wiki/SliceTricks
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
	if !reflect.DeepEqual(cutSlice(tests[1].slc1, 2), tests[1].exp) {
		panic(tests[1].slc1)
	}
	if !reflect.DeepEqual(expandSlice(tests[2].slc1, 2, 2), tests[2].exp) {
		panic(tests[2].slc1)
	}
	if !reflect.DeepEqual(extendSlice(tests[3].slc1, 2), tests[3].exp) {
		panic(tests[3].slc1)
	}
	reverseSlice(tests[4].slc1)
	if !reflect.DeepEqual(tests[4].slc1, tests[4].exp) {
		panic(tests[4].slc1)
	}
	l := len(tests[5].slc1)
	if !reflect.DeepEqual(insertSlice(tests[5].slc1, tests[5].slc2, l), tests[5].exp) {
		panic(tests[5].slc1)
	}
}

func insertSlice(source, insert []string, pos int) []string {
	l := len(source)
	if pos > l {
		return []string{}
	}
	source = append(source[:pos], append(insert, source[pos:]...)...)
	return source
}

func extendSlice(source []string, ext int) []string {
	source = append(source, make([]string, ext)...)
	return source
}

func expandSlice(source []string, pos, exp int) []string {
	// a = append(a[:i], append(make([]T, j), a[i:]...)...)
	l := len(source)
	if pos > l {
		return []string{}
	}
	source = append(source[:pos], append(make([]string, exp), source[pos:]...)...)
	return source
}

func cutSlice(source []string, pos int) []string {
	l := len(source)
	if pos >= l-1 {
		return []string{}
	}
	b := append(source[:pos], source[pos+1:]...)
	return b
}

func reverseSlice(source []string) {
	l := len(source)
	for i := l/2 - 1; i >= 0; i-- {
		opp := l - 1 - i
		source[i], source[opp] = source[opp], source[i]
	}
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
	tests = append(tests, test{[]string{"a", "b", "c", "d"}, []string{"d", "e", "f", "g"}, []string{"d", "c", "b", "a"}})
	tests = append(tests, test{[]string{"a", "b", "c", "d"}, []string{"e", "f", "g"}, []string{"a", "b", "c", "d", "e", "f", "g"}})
}
