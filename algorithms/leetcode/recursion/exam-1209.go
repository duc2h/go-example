package main

// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string-ii/

// You are given a string s and an integer k, a k duplicate removal consists of choosing k adjacent and equal letters from s and removing them, causing the left and the right side of the deleted substring to concatenate together.

// We repeatedly make k duplicate removals on s until we no longer can.

// Return the final string after all such duplicate removals have been made. It is guaranteed that the answer is unique.

// Example 1:

// Input: s = "abcd", k = 2
// Output: "abcd"
// Explanation: There's nothing to delete.
// Example 2:

// Input: s = "deeedbbcccbdaa", k = 3
// Output: "aa"
// Explanation:
// First delete "eee" and "ccc", get "ddbbbdaa"
// Then delete "bbb", get "dddaa"
// Finally delete "ddd", get "aa"
// Example 3:

// Input: s = "pbbcggttciiippooaais", k = 2
// Output: "ps"

// func main() {
// 	s := "pbbcggttciiippooaais"
// 	k := 2
// 	r := removeDuplicates(s, k)
// 	fmt.Println(r)
// }

func removeDuplicates(s string, k int) string {
	for i := 0; i <= len(s)-k; i++ {
		a := s[i]

		for x := 1; x < k; x++ {
			b := s[i+x]

			if a != b {
				break
			} else {
				if x+1 == k {
					a1 := s[:i]
					a2 := s[x+i+1:]

					s = a1 + a2

					return removeDuplicates(s, k)
				}
			}

		}
	}

	return s
}
