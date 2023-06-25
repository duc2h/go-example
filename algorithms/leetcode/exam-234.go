package main

// https://leetcode.com/problems/palindrome-linked-list/description/

// Given the head of a singly linked list, return true if it is a palindrome or false otherwise.

// example1:
// Input: head = [1,2,2,1]
// Output: true

// example2:
// Input: head = [1,2]
// Output: false

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// func main() {

// }

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome1(head *ListNode) bool {
	var (
		stack = []int{}
	)

	for n := head; n != nil; n = n.Next {
		stack = append(stack, n.Val)
	}

	left, right := 0, len(stack)-1

	for left < right {
		if stack[left] != stack[right] {
			return false
		}

		left++
		right--
	}

	return true
}
