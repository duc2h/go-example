package main

// https://leetcode.com/problems/missing-number/description/

// Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.

// Example 1:

// Input: nums = [3,0,1]
// Output: 2
// Explanation: n = 3 since there are 3 numbers, so all numbers are in the range [0,3]. 2 is the missing number in the range since it does not appear in nums.
// Example 2:

// Input: nums = [0,1]
// Output: 2
// Explanation: n = 2 since there are 2 numbers, so all numbers are in the range [0,2]. 2 is the missing number in the range since it does not appear in nums.
// Example 3:

// Input: nums = [9,6,4,2,3,5,7,0,1]
// Output: 8
// Explanation: n = 9 since there are 9 numbers, so all numbers are in the range [0,9]. 8 is the missing number in the range since it does not appear in nums.

// func main() {
// 	a := []int{3, 0, 1, 2}
// 	r := missingNumber2(a)
// 	fmt.Println(r)
// }

func missingNumber1(nums []int) int {
	total := 0
	temp := 0
	for i := 0; i < len(nums); i++ {
		total += i + 1
		temp += nums[i]
	}

	return total - temp
}

func missingNumber(nums []int) int {
	var (
		m = map[int]int{}
	)

	for _, v := range nums {
		m[v] = 0
	}

	for i := 0; i < len(nums)+1; i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}

	return -1
}

func missingNumber2(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	for i := 0; i < len(nums); i++ {
		if i != nums[i] {
			return i
		}
	}

	return len(nums)
}
