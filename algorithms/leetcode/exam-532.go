package main

import (
	"sort"
)

// https://leetcode.com/problems/k-diff-pairs-in-an-array/

// Given an array of integers nums and an integer k, return the number of unique k-diff pairs in the array.
// A k-diff pair is an integer pair (nums[i], nums[j]), where the following are true:
// 0 <= i, j < nums.length
// i != j
// |nums[i] - nums[j]| == k
// Notice that |val| denotes the absolute value of val.

// Example 1:
// Input: nums = [3,3,1,4,1,5], k = 2
// Output: 2
// Explanation: There are two 2-diff pairs in the array, (1, 3) and (3, 5).
// Although we have two 1s in the input, we should only return the number of unique pairs.

// Example 2:
// Input: nums = [1,2,3,4,5], k = 1
// Output: 4
// Explanation: There are four 1-diff pairs in the array, (1, 2), (2, 3), (3, 4) and (4, 5).

// Example 3:
// Input: nums = [1,3,1,5,4], k = 0
// Output: 1
// Explanation: There is one 0-diff pair in the array, (1, 1).

// Example 4:
// Input: nums = [3,3,1,4,1,5], k = -2
// Output: 2

// a := []int{1, 2, 4, 4, 3, 3, 0, 9, 2, 3}, k = 3 => expect = 2 , (3, 0), (4,1)
// func main() {
// 	a := []int{3, 3, 1, 4, 1, 5}
// 	r := findPairs1(a, 2)

// 	fmt.Println(r)
// }

func findPairs1(nums []int, k int) int {
	var (
		m      = map[int]int{}
		result = 0
	)

	sort.Ints(nums)

	for _, num := range nums {
		key := num - k
		if v, ok := m[key]; ok {
			if v == 0 {
				result++
				m[key]++
			}
		}

		if _, ok := m[num]; !ok {
			m[num] = 0
		}
	}

	return result
}

func findPairs(nums []int, k int) int {
	var (
		m      = map[int]int{}
		result = 0
	)

	// sort
	for i, _ := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]

			}
		}
	}

	for _, num := range nums {
		key := num - k

		if value, ok := m[key]; ok {
			// can be add or minus
			if value <= 1 {
				result++
				m[key]++
			}
		}

		if _, ok := m[num]; !ok {
			m[num] = 1
		}

	}

	return result
}
