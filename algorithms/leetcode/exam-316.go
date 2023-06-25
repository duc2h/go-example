package main

import "fmt"

// https://leetcode.com/problems/remove-duplicate-letters/

// Given a string s, remove duplicate letters so that every letter appears once and only once. You must make sure your result is
// the smallest in lexicographical order
//  among all possible results.

// Example 1:
// Input: s = "bcabc"
// Output: "abc"

// Example 2:
// Input: s = "cbacdcbc"
// Output: "acdb"

// b

//

// func main() {
// 	r := removeDuplicateLetters("cbacdcbc")

// 	fmt.Println(r)
// }

func removeDuplicateLetters(s string) string {

	var (
		mStr     = map[rune]int{}
		mVisited = map[rune]bool{}
		stack    = []rune{}
	)

	for _, k := range s {
		mStr[k]++
	}

	fmt.Println(mStr)

	for _, k := range s {

		if len(stack) > 0 {

			tail := stack[len(stack)-1]
			if tail >= k { // d c
				// remove tail, change visited
				// check mStr
				if v := mStr[tail]; v < 1 {
					if mVisited[k] {
						continue
					}

					stack = append(stack, k)
					continue

				}

				if mVisited[k] {
					continue
				}

				fmt.Println(string(k))
				fmt.Println(string(tail))

				stack = stack[:len(stack)-1]
				stack = append(stack, k)

			} else {
				// visited
				if mVisited[k] {
					continue
				}

				stack = append(stack, k)
				mVisited[tail] = true

			}

		} else {
			stack = append(stack, k)

		}

	}

	return string(stack)
}
