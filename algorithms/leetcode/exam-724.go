package main

// https://leetcode.com/problems/find-pivot-index/

// Given an array of integers nums, calculate the pivot index of this array.

// The pivot index is the index where the sum of all the numbers strictly to the left of the index is equal to the sum of all the numbers strictly to the index's right.

// If the index is on the left edge of the array, then the left sum is 0 because there are no elements to the left. This also applies to the right edge of the array.

// Return the leftmost pivot index. If no such index exists, return -1.

// Example 1:

// Input: nums = [1,7,3,6,5,6]
// Output: 3
// Explanation:
// The pivot index is 3.
// Left sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11
// Right sum = nums[4] + nums[5] = 5 + 6 = 11
// Example 2:

// Input: nums = [1,2,3]
// Output: -1
// Explanation:
// There is no index that satisfies the conditions in the problem statement.
// Example 3:

// Input: nums = [2,1,-1]
// Output: 0
// Explanation:
// The pivot index is 0.
// Left sum = 0 (no elements to the left of index 0)
// Right sum = nums[1] + nums[2] = 1 + -1 = 0

// func main() {
// 	a := []int{1, 2, 3}
// 	r := pivotIndex(a)
// 	fmt.Println(r)
// }

// 1,2,3,3 =>

// 8 - 2 = 6
func pivotIndex(nums []int) int {
	var total, sum, tmp int

	for _, v := range nums {
		total += v
	}

	for i, v := range nums {

		sum += tmp
		total = total - v

		if sum == total {
			return i
		}

		tmp = v
	}

	return -1
}

func pivotIndex1(nums []int) int {
	for i := 0; i < len(nums); i++ {
		var (
			left  = 0
			right = 0
		)

		for j := 0; j < i; j++ {
			left += nums[j]
		}

		for j := i + 1; j < len(nums); j++ {
			right += nums[j]
		}

		if left == right {
			return i
		}

	}

	return -1
}
