package main

// https://leetcode.com/problems/remove-element/

// func main() {
// 	a := []int{0, 1, 2, 2, 3, 0, 4, 2}
// 	fmt.Println(removeElement(a, 2))
// 	fmt.Println(a)
// }

func removeElement(nums []int, val int) int {
	index := 0
	for _, n := range nums {
		if n != val {
			nums[index] = n
			index++
		}
	}
	return index
}
