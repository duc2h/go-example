package main

// https://leetcode.com/problems/reverse-string/

// Write a function that reverses a string. The input string is given as an array of characters s.

// You must do this by modifying the input array in-place with O(1) extra memory.

// Example 1:
// Input: s = ['h','e','l','l','o']
// Output: ['o','l','l','e','h']

// Example 2:
// Input: s = ['H','a','n','n','a','h']
// Output: ['h','a','n','n','a','H']

// func main() {
// 	s := []byte{'H', 'a', '1', 'y', 'n', 'b', 'h'}
// 	reserveString(s)
// 	fmt.Println(string(s))
// }

func reserveString(s []byte) {
	lenth := len(s) / 2

	for i := 0; i < lenth; i++ {
		y := len(s) - (i + 1)

		v3 := s[i]
		s[i] = s[y]
		s[y] = v3
	}
}
