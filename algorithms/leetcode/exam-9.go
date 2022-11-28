package main

import (
	"strconv"
)

// https://leetcode.com/problems/palindrome-number/

// func main() {
// 	x := -1001

// 	fmt.Println(isPalindrome(x))
// }

func isPalindrome(x int) bool {

	s := strconv.Itoa(x)
	r := ""
	for i := len(s) - 1; i >= 0; i-- {
		r = r + string(s[i])
	}

	return s == r
}
