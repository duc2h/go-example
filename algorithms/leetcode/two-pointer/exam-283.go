package main

// https://leetcode.com/problems/move-zeroes/

// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

// Note that you must do this in-place without making a copy of the array.

// Example 1:
// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]

// Example 2:
// Input: nums = [0]
// Output: [0]

// func main() {
// 	a := []int{1, 0, 1, 0, 2, 0, 13}
// 	b := []int{0}
// 	moveZeroesBest1(a)
// 	moveZeroesBest1(b)
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

func moveZeroesBest1(nums []int) {
	index := 0

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num != 0 {
			nums[index] = num
			index++
		}
	}

	for i := index; i < len(nums); i++ {
		nums[i] = 0
	}

}

func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		for y := i + 1; y < len(nums); y++ {
			num1 := nums[i]
			num2 := nums[y]
			// swap
			if num1 == 0 && num2 != 0 {
				temp := num1
				nums[i] = num2
				nums[y] = temp
			}
		}
	}
}

func moveZeroesBest(nums []int) {
	index := 0
	for _, v := range nums {
		nums[index] = v
		if v != 0 {
			index++
		}
	}

	for i := index; i < len(nums); i++ {
		nums[i] = 0
	}
}
