package main

import (
	"math"
)

// https://leetcode.com/problems/reverse-integer/

// Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

// Example 1:

// Input: x = 123
// Output: 321
// Example 2:

// Input: x = -123
// Output: -321
// Example 3:

// Input: x = 120
// Output: 21

// func main() {

// 	s := reverse(123)

// 	fmt.Println(s)
// }

func reverse(x int) int {
	var (
		r          = 0
		isNegative = false
	)

	if x < 0 {
		isNegative = true
		x = -x
	}

	for x != 0 {
		s := x % 10
		x = x / 10

		r = r*10 + s
	}

	if isNegative {
		r = -r
	}

	if r > math.MaxInt32 || r < math.MinInt32 {
		return 0
	}

	return r
}
