package main

// https://leetcode.com/problems/remove-all-occurrences-of-a-substring/

// Given two strings s and part, perform the following operation on s until all occurrences of the
// substring part are removed:
// Find the leftmost occurrence of the substring part and remove it from s.
// Return s after removing all occurrences of part.
// A substring is a contiguous sequence of characters in a string.

// Example 1:
// Input: s = "daabcbaabcbc", part = "abc"
// Output: "dab"
// Explanation: The following operations are done:
// - s = "daabcbaabcbc", remove "abc" starting at index 2, so s = "dabaabcbc".
// - s = "dabaabcbc", remove "abc" starting at index 4, so s = "dababc".
// - s = "dababc", remove "abc" starting at index 3, so s = "dab".
// Now s has no occurrences of "abc".

// Example 2:
// Input: s = "axxxxyyyyb", part = "xy"
// Output: "ab"
// Explanation: The following operations are done:
// - s = "axxxxyyyyb", remove "xy" starting at index 4 so s = "axxxyyyb".
// - s = "axxxyyyb", remove "xy" starting at index 3 so s = "axxyyb".
// - s = "axxyyb", remove "xy" starting at index 2 so s = "axyb".
// - s = "axyb", remove "xy" starting at index 1 so s = "ab".
// Now s has no occurrences of "xy".

// func main() {
// 	a := removeOccurrences("axxxxyyyyb", "xy")

// 	fmt.Println(a)
// }

func removeOccurrences(s string, part string) string {
	if len(s) == 0 || len(part) == 0 {
		return s
	}

	for i := range s {
		p1 := part[0]
		s1 := s[i]

		if s1 == p1 {

			if i+len(part) > len(s) {
				return s
			}

			s2 := s[i : i+len(part)]
			if s2 == part {
				s11 := s[:i] + s[i+len(part):]
				return removeOccurrences(s11, part)
			}
		}
	}

	return s
}

// func removeOccurrences(s string, part string) string {
// 	if len(s) == 0 || len(part) == 0 {
// 		return part
// 	}

// 	for i := range part {
// 		s1 := s[0]
// 		p1 := part[i]

// 		if s1 == p1 {

// 			if i+len(s) > len(part) {
// 				return part
// 			}

// 			p2 := part[i : i+len(s)]
// 			if p2 == s {
// 				part1 := part[:i] + part[i+len(s):]
// 				fmt.Println(part1)

// 				return removeOccurrences(s, part1)
// 			}
// 		}
// 	}

// 	return part
// }
