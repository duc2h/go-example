package main

import "fmt"

// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/

func main() {

	// fmt.Println(strStr("sabutsad", "sad"))

	fmt.Println(strStr("mississippi", "issip"))
}

func strStr(haystack string, needle string) int {
	index := 0
	tmp := 0
	for i := 0; i < len(haystack); i++ {
		s := haystack[i]

		if s == needle[index] {
			if index == 0 {
				tmp = i
			}
			index++
			if index == len(needle) {
				return tmp
			}
		} else {
			index = 0
			if s == needle[index] {
				tmp = i
				index++
			}
		}
	}

	return -1
}
