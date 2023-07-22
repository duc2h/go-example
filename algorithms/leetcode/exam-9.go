package main

import (
	"strconv"
)

// https://leetcode.com/problems/palindrome-number/

// Given an integer x, return true if x is a
// palindrome
// , and false otherwise.

// Example 1:
// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.

// Example 2:
// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

// Example 3:
// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

// func main() {
// 	x := 1001

// 	fmt.Println(isPalindrome11(x))
// }

func isPalindrome11(x int) bool {
	r := 0
	x1 := x

	if x < 0 {
		return false
	}

	for x != 0 {
		s := x % 10
		x = x / 10

		r = r*10 + s
	}

	return r == x1

}

func isPalindrome(x int) bool {

	s := strconv.Itoa(x)
	r := ""
	for i := len(s) - 1; i >= 0; i-- {
		r = r + string(s[i])
	}

	return s == r
}
