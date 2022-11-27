package main

// https://leetcode.com/problems/valid-parentheses/

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// 1. Open brackets must be closed by the same type of brackets.
// 2. Open brackets must be closed in the correct order.
// 3. Every close bracket has a corresponding open bracket of the same type.

// Example 1:
// Input: s = "()"
// Output: true

// Example 2:
// Input: s = "()[]{}"
// Output: true

// Example 3:
// Input: s = "(]"
// Output: false

// func main() {
// 	fmt.Println(isValid("[][]{}"))
// 	fmt.Println(isValid("[][(]){}"))
// }

func isValid(s string) bool {
	length := len(s)
	if length == 0 || length%2 != 0 {
		return false
	}

	m := map[rune]rune{}
	m['{'] = '}'
	m['('] = ')'
	m['['] = ']'
	stack := []rune{}
	for _, k := range s {
		if v, ok := m[k]; ok {
			stack = append(stack, v)
		} else if len(stack) == 0 || stack[len(stack)-1] != k {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
