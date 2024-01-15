package main

import "fmt"

// https://leetcode.com/problems/container-with-most-water/description/

// You are given an integer array height of length n. There are n vertical lines drawn such
// that the two endpoints of the ith line are (i, 0) and (i, height[i]).

// Find two lines that together with the x-axis form a container, such that the container contains the most water.
// Return the maximum amount of water a container can store.
// Notice that you may not slant the container.

// Input: height = [1,8,6,2,5,4,8,3,7]
// Output: 49
// Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
// Example 2:

// Input: height = [1,1]
// Output: 1

func main() {
	a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	r := maxArea(a)
	fmt.Println(r)
}

func maxArea(nums []int) int {
	max := 0

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			x := j - i
			sub := nums[i] - nums[j]
			max_local := 0
			if sub == 0 {
				max_local = x * nums[i]
			} else {
				if nums[i] > nums[j] {
					max_local = x * nums[j]
				} else {
					max_local = x * nums[i]
				}
			}

			if max_local < 0 {
				max_local = -max_local
			}

			if max_local > max {
				max = max_local
			}
		}
	}

	return max
}
