package main

import (
	"fmt"
	"strconv"
)

// Given an encoded string, return its decoded string.
// The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times. Note that k is guaranteed to be a positive integer.
// You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed, etc. Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k. For example, there will not be input like 3a or 2[4].
// The test cases are generated so that the length of the output will never exceed 105.

// Example 1:
// Input: s = "3[a]2[bc]"
// Output: "aaabcbc"

// Example 2:
// Input: s = "3[a2[c]]"
// Output: "accaccacc"

// Example 3:
// Input: s = "2[abc]3[cd]ef"
// Output: "abcabccdcdcdef"

// define string in [] and number front of [].

// func main() {

// 	r := decodeString("21[ba2[c]]")
// 	fmt.Println(r)
// }

// stack = 2, [, a, b

// ba -> baba

type Stack []string

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() string {
	index := len(*s) - 1
	v := (*s)[index]

	*s = (*s)[:index]

	return v
}

func (s *Stack) PushLongStr(str string) {
	for i := len(str) - 1; i >= 0; i-- {
		*s = append(*s, string(str[i]))
	}
}

func decodeString(s string) string {
	stack := Stack{}

	for _, v := range s {
		if v == ']' {
			// fmt.Println(stack)
			str := ""
			for len(stack) > 0 {
				e := stack.Pop()
				if e == "[" {
					break
				}

				str += e
			}

			numStr := ""
			for len(stack) > 0 {
				// fmt.Println(stack)
				t := stack.Pop()
				// fmt.Println(t)

				if _, err := strconv.Atoi(t); err == nil {
					numStr = t + numStr

				} else {
					stack.Push(t)
					break
				}
			}

			str1 := ""
			if num, err := strconv.Atoi(numStr); err == nil {
				// fmt.Println(num)
				for i := 0; i < num; i++ {
					str1 += str

				}
			}

			// fmt.Println(str1)
			// fmt.Println(str)
			fmt.Println(stack)

			stack.PushLongStr(str1)

		} else {
			stack.Push(string(v))
		}

	}

	str := ""

	for _, v := range stack {
		str += v
	}

	return str
}
