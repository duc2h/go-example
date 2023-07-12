package main

// https://leetcode.com/discuss/interview-question/algorithms/124938/google-subarray-sum-equals-target

// Given a non-empty array nums contains positive integers and a positive integer target.
// Find the first subarray in nums that sums up to target and return the begin and end index of this subarray. If there is no such subarray, return [-1, -1].

// Example 1:
// Input: nums = [4, 3, 5, 7, 8], target = 12
// Output: [0, 2]
// Explanation: 4 + 3 + 5 = 12. Although 5 + 7 = 12, [4, 3, 5] is the first subarray that sums up to 12.

// Example 2:
// Input: nums = [1, 2, 3, 4], target = 15
// Output: [-1, -1]
// Explanation: Thers is no such subarray that sums up to 15.

// Example 3:
// Input: nums = [4,1,2,3, 0], k = 3
// Output: [1,2]

// Example 4:
// Input: nums = [0,2,3], k = 3
// Output: [2,2]

// func main() {
// 	nums := []int{2, 3, 0}
// 	r := findSubArray(nums, 3)

// 	fmt.Println(r)
// }

func findSubArray(nums []int, k int) []int {
	var (
		result = []int{-1, -1}
		total  = 0
		j      = 0
	)

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		total += num

		for total > k {
			total = total - nums[j]
			j++
		}

		if total == k {
			result[0] = j
			result[1] = i
			break
		}

	}

	return result
}
