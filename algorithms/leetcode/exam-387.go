package main

import "fmt"

// https://leetcode.com/problems/first-unique-character-in-a-string/description/

// func main() {

// 	s := "leetcode"
// 	a := firstUniqChar(s)
// 	fmt.Println(a)

// }

// Input: s = "leetcode"
// Output: 0
// Example 2:

// Input: s = "loveleetcode"
// Output: 2
// Example 3:

// Input: s = "aabb"
// Output: -1

func firstUniqChar(s string) int {
	m := make(map[rune]int)

	for _, v := range s {
		m[v]++
	}

	fmt.Println(m)
	for i, k := range s {

		if v := m[k]; v == 1 {

			return i
		}

	}

	return -1
}
