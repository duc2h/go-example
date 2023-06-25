package main

import (
	"sort"
)

// https://leetcode.com/problems/3sum/

// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
// Notice that the solution set must not contain duplicate triplets.

// Example 1:
// Input: nums = [-1,0,1,2,-1,-4]
// Output: [[-1,-1,2],[-1,0,1]]
// Explanation:
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
// The distinct triplets are [-1,0,1] and [-1,-1,2].
// Notice that the order of the output and the order of the triplets does not matter.

// Example 2:
// Input: nums = [0,1,1]
// Output: []
// Explanation: The only possible triplet does not sum up to 0.

// []int{0, 0, 0, 0} => [[0,0,0]]

// func main() {

// 	a := []int{1, 0, -1, 1, 0, -1}
// 	r := threeSum(a)

// 	fmt.Println(r)
// }

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var (
		result = [][]int{}
	)

	for i := 0; i < len(nums)-2; i++ {
		l := i + 1
		r := len(nums) - 1

		for l < r {
			total := nums[i] + nums[l] + nums[r]

			if total == 0 {
				a := []int{nums[i], nums[l], nums[r]}
				result = append(result, a)
				break
			} else if total < 0 {
				l++
			} else {
				r--
			}

		}
	}

	return result
}
