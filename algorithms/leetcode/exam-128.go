package main

// https://leetcode.com/problems/longest-consecutive-sequence/

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
	smallest := nums[0]
	for _, num := range nums {
		m[num] = true
		if smallest > num {
			smallest = num
		}
	}

	// find the smallest, then increase 1 in map.
	for _, num := range nums {
		if m[num-1] {
			continue
		}

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
