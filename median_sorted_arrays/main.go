// Leetcode problem
//There are two sorted arrays nums1 and nums2 of size m and n respectively.
//
//Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
//
//You may assume nums1 and nums2 cannot be both empty.
//
//Example 1:
//
//nums1 = [1, 3]
//nums2 = [2]
//
//The median is 2.0
//Example 2:
//
//nums1 = [1, 2]
//nums2 = [3, 4]
//
//The median is (2 + 3)/2 = 2.5

// This solution does not use extra buffers (merged slice)
// and does everything in a single pass O(m+n)

package main

func main() {
	findMedianSortedArrays([]int{1, 2}, []int{3, 4})
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	l1 := len(nums1)
	l2 := len(nums2)
	l := l1 + l2
	var even, n1even, n2even bool
	var m, m1, m2 float64
	h := l / 2

	if l%2 == 0 {
		even = true
	}
	if l1%2 == 0 {
		n1even = true
	}
	if l2%2 == 0 {
		n2even = true
	}

	if l1 == 0 {
		if n2even {
			m = (float64(nums2[h]) + float64(nums2[h-1])) / 2
		} else {
			m = float64(nums2[h])
		}
		return m
	}

	if l2 == 0 {
		if n1even {
			m = (float64(nums1[h]) + float64(nums1[h-1])) / 2
		} else {
			m = float64(nums1[h])
		}
		return m
	}

	n1idx, n2idx := 0, 0

	for i := 0; i <= h; i++ {
		if n1idx < l1 && n2idx < l2 {
			if nums1[n1idx] <= nums2[n2idx] {
				if i == h-1 {
					m1 = float64(nums1[n1idx])
				}
				if i == h {
					m2 = float64(nums1[n1idx])
				}
				n1idx++
			} else {
				if i == h-1 {
					m1 = float64(nums2[n2idx])
				}
				if i == h {
					m2 = float64(nums2[n2idx])
				}
				n2idx++
			}
		} else if n1idx < l1 {
			if i == h-1 {
				m1 = float64(nums1[n1idx])
			}
			if i == h {
				m2 = float64(nums1[n1idx])
			}
			n1idx++
		} else if n2idx < l2 {
			if i == h-1 {
				m1 = float64(nums2[n2idx])
			}
			if i == h {
				m2 = float64(nums2[n2idx])
			}
			n2idx++
		} else {
			panic("barf!")
		}
	}
	if even {
		m = (m1 + m2) / 2
	} else {
		m = m2
	}
	return m
}
