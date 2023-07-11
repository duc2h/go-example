package main

import "fmt"

// Given a string s, return the longest
// palindromic substring in s.

// Example 1:
// Input: s = "babad"
// Output: "bab"
// Explanation: "aba" is also a valid answer.

// Example 2:
// Input: s = "cbbd"
// Output: "bb"

func main() {
	s := longestPalindrome("dbbd")
	fmt.Println(s)
}

// babad
// step1: loop1 l: 0, r: 0, len: 0 => res: b,  l: -1, r: 1, len: 1
// step1: loop2 l:0, r:1, len: 1 =>  ''
// step2: loop1 l: 1, r: 1, len: 1 => res: a, l:0, r:2, len: 1
// step2: loop1 l: 0, r: 2, len: 1 => res: bab, l: -1, r: 3, len: 3
// step2: loop2 l: 1, r: 2, len: 3 =>  ""
// step3: loop1 l: 2, r: 2, len: 3 =>   ''

// cbbd
// step1: loop1: l: 0, r: 0, len: 0 => res: c,  l: -1, r: 1, len: 1
// step1: loop2: l: 0, r: 1, len: 1 => res: c, l: -1, r: 2, len: 2
// step2: loop1: l: 1, r: 1, len: 2 => res: c, l: 0, r: 2, len: 2
// step2: loop1: l: 0, r:

func longestPalindrome(s string) string {

	var (
		res    = ""
		length = 0
	)

	for i := range s {

		// odd length
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if (r - l + 1) > length {
				res = s[l : r+1]
				length = r - l + 1
			}

			l--
			r++
		}

		// even length
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if (r - l + 1) > length {
				res = s[l : r+1]
				length = r - l + 1
			}

			l--
			r++

		}

	}

	return res
}
