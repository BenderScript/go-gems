// Leetcode problem
//Given a string, find the longest substring without repeating characters.
//
//Example 1:
//
//Input: "abcabcbb"
//Output: 3
//Explanation: The answer is "abc", with the length of 3.
//Example 2:
//
//Input: "bbbbb"
//Output: 1
//Explanation: The answer is "b", with the length of 1.
//Example 3:
//
//Input: "pwwkew"
//Output: 3
//Explanation: The answer is "wke", with the length of 3.
//Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

package main

type test struct {
	str1 string
	exp  string
}

var tests []test

func main() {
	for _, t := range tests {
		if longestSubstring(t.str1) != t.exp {
			panic("Test failed:" + t.str1)
		}
	}
}

func longestSubstring(s string) string {
	l := len(s)
	sub := ""
	subMax := ""
	subS := 0
	lSub, lMax := 0, 0
	cMap := make(map[uint8]int, 0)
	for i := 0; i < l; i++ {
		ch := s[i]
		if v, ok := cMap[ch]; ok {
			if lSub > lMax {
				lMax = lSub
				subMax = sub
			}
			sub = s[v+1 : i+1]
			lSub = len(sub)
			cMap[ch] = i
			for j := subS; j < v; j++ {
				delete(cMap, s[j])
			}
			subS = v + 1
		} else {
			cMap[ch] = i
			sub += string(ch)
			lSub++
		}
	}
	if lSub > lMax {
		lMax = lSub
		subMax = sub
	}
	return subMax
}

func init() {
	tests = make([]test, 0)
	tests = append(tests, test{" ", " "})
	tests = append(tests, test{"abcabcbb", "abc"})
	tests = append(tests, test{"bbbb", "b"})
	tests = append(tests, test{"pwwkew", "wke"})
}
