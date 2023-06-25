package main

// https://leetcode.com/problems/binary-search/

// Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums.
// If target exists, then return its index. Otherwise, return -1.

// You must write an algorithm with O(log n) runtime complexity.

// Example 1:

// Input: nums = [-1,0,3,5,9,12], target = 9
// Output: 4
// Explanation: 9 exists in nums and its index is 4
// Example 2:

// Input: nums = [-1,0,3,5,9,12], target = 2
// Output: -1
// Explanation: 2 does not exist in nums so return -1

// func main() {
// 	a := []int{-1, 0, 3, 5, 9, 12}
// 	r := search(a, 9)
// 	fmt.Println(r)
// }

func search(nums []int, target int) int {

	if len(nums) == 0 {
		return -1
	}
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, left, right, target int) int {
	if right >= left {

		mid := (left + right) / 2
		num := nums[mid]
		if num == target {
			return mid
		} else if num > target {
			return binarySearch(nums, left, mid-1, target)
		} else {
			return binarySearch(nums, mid+1, right, target)
		}
	}
	return -1
}
