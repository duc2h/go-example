package main

// https://leetcode.com/problems/rotate-array/description/

// Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

// Example 1:

// Input: nums = [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]
// Explanation:
// rotate 1 steps to the right: [7,1,2,3,4,5,6]
// rotate 2 steps to the right: [6,7,1,2,3,4,5]
// rotate 3 steps to the right: [5,6,7,1,2,3,4]
// Example 2:

// Input: nums = [-1,-100,3,99], k = 2
// Output: [3,99,-1,-100]
// Explanation:
// rotate 1 steps to the right: [99,-1,-100,3]
// rotate 2 steps to the right: [3,99,-1,-100]

// len = 4
// k = 2
// index = 2 => index = 4, 4 - len = 0

// func main() {

// 	a := []int{1, 2, 3, 4, 5, 6, 7}
// 	rotate(a, 3)

// 	fmt.Println(a)
// }

func rotate(nums []int, k int) {

	for i := 0; i < k; i++ {
		previous := nums[len(nums)-1]

		for j := 0; j < len(nums); j++ {
			nums[j], previous = previous, nums[j]
		}

	}

}
