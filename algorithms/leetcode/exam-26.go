package main

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

// func main() {
// 	a := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
// 	// b := []int{1, 1, 2}
// 	fmt.Println(removeDuplicates(a))
// 	fmt.Println(a)
// }

func removeDuplicates(nums []int) int {
	index := 0
	for _, n := range nums {
		if n != nums[index] {
			index++
			nums[index] = n
		}
	}

	return index + 1
}
