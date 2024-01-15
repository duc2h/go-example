package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/word-break/

// Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.
// Note that the same word in the dictionary may be reused multiple times in the segmentation.

// Example 1:
// Input: s = "leetcode", wordDict = ["leet","code"]
// Output: true
// Explanation: Return true because "leetcode" can be segmented as "leet code".

// Example 2:
// Input: s = "applepenapple", wordDict = ["apple","pen"]
// Output: true
// Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
// Note that you are allowed to reuse a dictionary word.

// Example 3:
// Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
// Output: false

// s = "bb"
// wordDict = ["a","b","bbb","bbbb"]

func main() {

	a := []string{"cats", "dog", "sand", "cat"}
	r := wordBreak("catsanddog", a)
	fmt.Println(r)
}

func wordBreak(s string, wordDict []string) bool {
	arr := []string{s}
	mapDuplicate := map[string]bool{}
	for len(arr) > 0 {
		// dequeue()
		remaining := arr[0]
		arr = arr[1:]
		if mapDuplicate[remaining] {
			continue
		}
		for _, word := range wordDict {
			if word == remaining {
				return true
			}

			if strings.HasPrefix(remaining, word) {
				arr = append(arr, remaining[len(word):])
				mapDuplicate[remaining] = true
			}
		}
	}

	return false
}
