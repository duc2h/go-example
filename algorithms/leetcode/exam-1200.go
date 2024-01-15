package main

import (
	"math"
)

// https://leetcode.com/problems/minimum-absolute-difference/

// Given an array of distinct integers arr, find all pairs of elements with the minimum
// absolute difference of any two elements.

// Return a list of pairs in ascending order(with respect to pairs), each pair [a, b] follows

// a, b are from arr
// a < b
// b - a equals to the minimum absolute difference of any two elements in arr

// Example 1:

// Input: arr = [4,2,1,3]
// Output: [[1,2],[2,3],[3,4]]
// Explanation: The minimum absolute difference is 1. List all pairs with difference equal to 1 in ascending order.
// Example 2:

// Input: arr = [1,3,6,10,15]
// Output: [[1,3]]
// Example 3:

// Input: arr = [3,8,-10,23,19,-4,-14,27]
// Output: [[-14,-10],[19,23],[23,27]]

// func main() {
// 	arr := []int{3, 8, -10, 23, 19, -4, -14, 27}
// 	r := minimumAbsDifference(arr)

// 	fmt.Println(r)
// }

func minimumAbsDifference(arr []int) [][]int {

	result := [][]int{}
	mNum := map[int]int{}
	minNum := math.MaxInt

	for _, v := range arr {
		mNum[v] = 0
	}
	// sort
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	// find min
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j]-arr[i] < minNum {
				minNum = arr[j] - arr[i]
			}
		}
	}

	for _, v := range arr {
		if _, exist := mNum[v+minNum]; exist {
			result = append(result, []int{v, v + minNum})
		}
	}

	return result
}
