package main

// cannot apply sliding window.
// https://leetcode.com/problems/subarray-sum-equals-k/

// Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.
// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [1,1,1], k = 2
// Output: 2 => [1,1] [1,1]

// Example 2:
// Input: nums = [1,2,3], k = 3
// Output: 2, => [1,2], [3]

// Example 3:
// Input: nums = [4,1,2,3, 0], k = 3
// Output: 3, => [1,2], [3], [3, 0]

// j= 1, i = 2

// func main() {
// 	nums := []int{-1, -1, 1}
// 	r := subarraySum(nums, 0)
// 	fmt.Println(r)
// }

// func subarraySum(nums []int, k int) int {

// }
