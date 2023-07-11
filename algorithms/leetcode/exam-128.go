package main

// https://leetcode.com/problems/longest-consecutive-sequence/

// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

// You must write an algorithm that runs in O(n) time.

// Example 1:

// Input: nums = [100,4,200,1,3,2]
// Output: 4
// Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
// Example 2:

// Input: nums = [0,3,7,2,5,8,4,6,0,1]
// Output: 9

// func main() {
// 	a := []int{100, 4, 200, 1, 3, 2}
// 	fmt.Println(longestConsecutive(a))
// }

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := 0
	m := map[int]bool{}
	// smallest := nums[0]
	for _, num := range nums {
		m[num] = true
		// if smallest > num {
		// 	smallest = num
		// }
	}

	// find the smallest, then increase 1 in map.
	for _, num := range nums {
		// if m[num-1] {
		// 	continue
		// }

		count := 1
		for m[num+1] {
			count++
			num++
		}
		if count > max {
			max = count
		}
	}

	return max
}
