package main

// https://leetcode.com/problems/check-if-the-sentence-is-pangram/description/

// A pangram is a sentence where every letter of the English alphabet appears at least once.
// Given a string sentence containing only lowercase English letters, return true if sentence is a pangram, or false otherwise.

// Example 1:
// Input: sentence = "thequickbrownfoxjumpsoverthelazydog"
// Output: true
// Explanation: sentence contains at least one of every letter of the English alphabet.

// Example 2:
// Input: sentence = "leetcode"
// Output: false

// func main() {
// 	r := checkIfPangram("leetcode")
// 	fmt.Println(r)
// }

func checkIfPangram(sentence string) bool {
	var (
		m     = map[rune]bool{}
		count = 0
		str   = "qwertyuioplkjhgfdsazxcvbnm"
	)

	for _, s := range str {
		m[s] = false
	}

	for _, k := range sentence {
		if v, ok := m[k]; ok && !v {
			m[k] = true
			count++
		}
	}

	return count == 26
}
