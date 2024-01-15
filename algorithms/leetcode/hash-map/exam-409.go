package main

import "fmt"

// https://leetcode.com/problems/longest-palindrome/description/

// Given a string s which consists of lowercase or uppercase letters, return the length of the longest palindrome that can be built with those letters.

// Letters are case sensitive, for example, "Aa" is not considered a palindrome here.

// Example 1:

// Input: s = "abcccQQcdd"
// Output: 7
// Explanation: One longest palindrome that can be built is "dccaccd", whose length is 7.
// Example 2:

// Input: s = "a"
// Output: 1
// Explanation: The longest palindrome that can be built is "a", whose length is 1.

func main() {
	r := longestPalindrome("ababababa")
	fmt.Println(r)
}

func longestPalindrome(s string) int {
	m := map[rune]int{}
	result := 0
	flag := false
	for _, k := range s {
		m[k]++
	}

	for _, v := range m {
		if v%2 == 0 {
			result += v
		} else {
			result += v - 1
			flag = true
		}
	}

	if flag {
		result++
	}

	return result
}
