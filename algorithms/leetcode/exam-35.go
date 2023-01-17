package main

import "fmt"

// https://leetcode.com/problems/search-insert-position/

func main() {
	// a := []int{1, 3, 5, 6}
	a := []int{1}
	fmt.Println(searchInsert(a, 0))
}

func searchInsert(nums []int, target int) int {
	return binarySearch1(nums, target, 0, len(nums)-1)
}

func binarySearch1(nums []int, target, left, right int) int {

	mid := (left + right) / 2
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		if mid-1 == left {
			if target < nums[left] {
				return left
			} else {
				return mid
			}
		}

		return binarySearch1(nums, target, left, mid-1)
	} else {
		fmt.Println("mid: ", mid)
		fmt.Println("right: ", right)

		if mid+1 == right {
			if target < nums[right] {
				fmt.Println("right: ", right)

				return right
			} else {
				fmt.Println("right: ", right)

				return right + 1
			}
		}

		return binarySearch1(nums, target, mid+1, right)

	}
}
