package main

// still fail

// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/

// func main() {

// 	// fmt.Println(strStr("sabutsad", "sad"))

// 	fmt.Println(strStr("mississippi", "issip"))
// }

func strStr(haystack string, needle string) int {
	length := len(needle)
	for i := 0; i < len(haystack); i++ {
		index := i + length
		if index > len(haystack) {
			return -1
		}

		sub := haystack[i:index]

		if sub == needle {
			return i
		}
	}

	return -1
}
