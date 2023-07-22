package main

import "fmt"

// https://leetcode.com/problems/two-out-of-three/

// Given three integer arrays nums1, nums2, and nums3, return a distinct array containing all the values that are present in at least two out of the three arrays.
// You may return the values in any order.

// Example 1:

// Input: nums1 = [1,1,3,2], nums2 = [2,3], nums3 = [3]
// Output: [3,2]
// Explanation: The values that are present in at least two arrays are:
// - 3, in all three arrays.
// - 2, in nums1 and nums2.
// Example 2:

// Input: nums1 = [3,1], nums2 = [2,3], nums3 = [1,2]
// Output: [2,3,1]
// Explanation: The values that are present in at least two arrays are:
// - 2, in nums2 and nums3.
// - 3, in nums1 and nums2.
// - 1, in nums1 and nums3.

// func main() {
// 	n1 := []int{2, 3}
// 	n2 := []int{1, 1, 3, 2}
// 	n3 := []int{3}
// 	r := twoOutOfThree(n1, n2, n3)
// 	fmt.Println(r)
// }

func uniqueArr(nums []int) []int {
	var (
		m      = map[int]bool{}
		result = make([]int, 0, len(nums))
	)

	for _, k := range nums {
		if !m[k] {
			m[k] = true

			result = append(result, k)
		}
	}

	return result
}

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	var (
		m      = map[int]int{}
		result = []int{}
	)

	for _, v := range uniqueArr(nums1) {
		m[v]++
	}

	for _, v := range uniqueArr(nums2) {
		m[v]++
	}

	for _, v := range uniqueArr(nums3) {
		m[v]++
	}

	fmt.Println(m)

	for k, v := range m {
		if v >= 2 {
			result = append(result, k)
		}
	}

	return result
}
