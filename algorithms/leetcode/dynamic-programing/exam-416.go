package main

// https://leetcode.com/problems/partition-equal-subset-sum/

// Given an integer array nums, return true if you can partition the array into two subsets
// such that the sum of the elements in both subsets is equal or false otherwise.

// Example 1:
// Input: nums = [1,5,11,5]
// Output: true
// Explanation: The array can be partitioned as [1, 5, 5] and [11].

// Example 2:
// Input: nums = [1,2,3,6]
// Output: false
// Explanation: The array cannot be partitioned into equal sum subsets.

// func main() {
// 	nums := []int{1, 2, 3, 5}
// 	r := canPartition(nums)

// 	fmt.Println(r)
// }

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true
	for _, v := range nums {
		for i := target; i >= 0; i-- {
			if dp[i] == true && i+v <= target {
				dp[v+i] = true
			}

			if dp[target] == true {
				return true
			}
		}
	}

	return false
}
