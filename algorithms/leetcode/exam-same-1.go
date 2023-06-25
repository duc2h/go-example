package main

// https://leetcode.com/discuss/interview-question/372434

// Given an int array nums and an int target, find how many unique pairs in the array such that their sum is equal to target. Return the number of pairs.

// Example 1:

// Input: nums = [1, 1, 2, 45, 46, 46], target = 47
// Output: 2
// Explanation:
// 1 + 46 = 47
// 2 + 45 = 47
// Example 2:

// Input: nums = [1, 1], target = 2
// Output: 1
// Explanation:
// 1 + 1 = 2
// Example 3:

// Input: nums = [1, 5, 1, 5], target = 6
// Output: 1
// Explanation:
// [1, 5] and [5, 1] are considered the same.

// example 3
// [4,-2,4,-2,4,-2], target = 2

// func main() {

// 	a := []int{2, 1, 2, 7, 8, 7, 11, 15}
// 	r := twoSum1(a, 9)

// 	fmt.Println(r)
// }

func twoSum1(nums []int, target int) int {
	var (
		m      = map[int]int{}
		result = 0
	)

	for _, v := range nums {
		k := target - v

		if v1, ok := m[k]; ok {
			if v1 == 1 {
				result++
				m[k]++
			}
		} else {
			if _, ok := m[v]; !ok {
				m[v] = 1
			}
		}
	}

	return result
}
