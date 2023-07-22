package main

import "fmt"

// https://leetcode.com/problems/search-insert-position/

// Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

// You must write an algorithm with O(log n) runtime complexity.

// Example 1:

// Input: nums = [1,3,5,6], target = 5
// Output: 2
// Example 2:

// Input: nums = [1,3,5,6], target = 2
// Output: 1
// Example 3:

// Input: nums = [1,3,5,6], target = 7
// Output: 4

// func main() {
// 	// a := []int{1, 3, 5, 6}
// 	a := []int{1}
// 	fmt.Println(searchInsert(a, 0))
// }

func searchInsert(nums []int, target int) int {
	return binarySearch1(nums, target, 0, len(nums)-1)
}

func binarySearch1(nums []int, target, left, right int) int {

	mid := (left + right) / 2
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		if mid-1 == left {
			if target < nums[left] {
				return left
			} else {
				return mid
			}
		}

		return binarySearch1(nums, target, left, mid-1)
	} else {
		fmt.Println("mid: ", mid)
		fmt.Println("right: ", right)

		if mid+1 == right {
			if target < nums[right] {
				fmt.Println("right: ", right)

				return right
			} else {
				fmt.Println("right: ", right)

				return right + 1
			}
		}

		return binarySearch1(nums, target, mid+1, right)

	}
}
