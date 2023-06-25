package main

// https://leetcode.com/problems/set-mismatch/

// You have a set of integers s, which originally contains all the numbers from 1 to n. Unfortunately, due to some error, one of the numbers in s
// got duplicated to another number in the set, which results in repetition of one number and loss of another number.
// You are given an integer array nums representing the data status of this set after the error.
// Find the number that occurs twice and the number that is missing and return them in the form of an array.

// Example 1:

// Input: nums = [1,2,2,4]
// Output: [2,3]
// Example 2:

// Input: nums = [1,1]
// Output: [1,2]

// func main() {

// 	a := []int{1, 1}
// 	r := findErrorNums(a)
// 	fmt.Println(r)
// }

func findErrorNums(nums []int) []int {
	var (
		m      = map[int]bool{}
		result = []int{}
	)

	for _, k := range nums {
		if m[k] {
			result = append(result, k)
		} else {
			m[k] = true
		}
	}

	for i := 1; i < len(nums)+1; i++ {
		if !m[i] {
			result = append(result, i)
			break
		}
	}

	return result
}
