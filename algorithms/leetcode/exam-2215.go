package main

// https://leetcode.com/problems/find-the-difference-of-two-arrays/

// Given two 0-indexed integer arrays nums1 and nums2, return a list answer of size 2 where:
// answer[0] is a list of all distinct integers in nums1 which are not present in nums2.
// answer[1] is a list of all distinct integers in nums2 which are not present in nums1.
// Note that the integers in the lists may be returned in any order.

// Example 1:

// Input: nums1 = [1,2,3], nums2 = [2,4,6]
// Output: [[1,3],[4,6]]
// Explanation:
// For nums1, nums1[1] = 2 is present at index 0 of nums2, whereas nums1[0] = 1 and nums1[2] = 3 are not present in nums2. Therefore, answer[0] = [1,3].
// For nums2, nums2[0] = 2 is present at index 1 of nums1, whereas nums2[1] = 4 and nums2[2] = 6 are not present in nums2. Therefore, answer[1] = [4,6].
// Example 2:

// Input: nums1 = [1,2,3,3], nums2 = [1,1,2,2]
// Output: [[3],[]]
// Explanation:
// For nums1, nums1[2] and nums1[3] are not present in nums2. Since nums1[2] == nums1[3], their value is only included once and answer[0] = [3].
// Every integer in nums2 is present in nums1. Therefore, answer[1] = [].

// func main() {
// 	n1 := []int{1, 2, 3, 3}
// 	n2 := []int{1, 1, 2, 2}
// 	r := findDifference(n1, n2)

// 	fmt.Println(r)
// }

func findDifference(nums1 []int, nums2 []int) [][]int {

	return [][]int{findDiff(nums1, nums2), findDiff(nums2, nums1)}
}

func findDiff(nums1 []int, nums2 []int) []int {
	var (
		m      = map[int]bool{}
		result = []int{}
	)

	for _, k := range nums2 {
		m[k] = true
	}

	for _, k := range nums1 {
		if !m[k] {
			result = append(result, k)
			m[k] = true
		}
	}

	return result
}