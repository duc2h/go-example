package main

// https://leetcode.com/problems/minimum-size-subarray-sum/

// Given an array of positive integers nums and a positive integer target, return the minimal length of a
// subarray
//  whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.

// Example 1:
// Input: target = 7, nums = [2,3,1,2,4,3]
// Output: 2
// Explanation: The subarray [4,3] has the minimal length under the problem constraint.

// Example 2:
// Input: target = 4, nums = [1,4,4]
// Output: 1

// Example 3:
// Input: target = 11, nums = [1,1,1,1,1,1,1,1]
// Output: 0

// func main() {
// 	a := []int{2, 3, 1, 2, 4, 3}
// 	r := minSubArrayLen(7, a)
// 	fmt.Println(r)
// }

func minSubArrayLen(target int, nums []int) int {
	var (
		total = 0
		min   = 0
		j     = 0
	)

	for i := 0; i < len(nums); i++ {
		total += nums[i]

		for total >= target {
			a := nums[j : i+1]
			min = finMin(min, len(a))

			total = total - nums[j]
			j++
		}

	}

	return min
}

func finMin(min, num2 int) int {
	if min != 0 && min < num2 {
		return min
	}

	return num2
}
