package main

// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and
// removing all non-alphanumeric characters, it reads the same forward and backward.
// Alphanumeric characters include letters and numbers.

// Given a string s, return true if it is a palindrome, or false otherwise.

// Example 1:
// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.

// Example 2:
// Input: s = "race a car"
// Output: false
// Explanation: "raceacar" is not a palindrome.

// Example 3:
// Input: s = " "
// Output: true
// Explanation: s is an empty string "" after removing non-alphanumeric characters.
// Since an empty string reads the same forward and backward, it is a palindrome.

// func main() {
// 	s := "aa"
// 	r := isPalindrome(s)
// 	fmt.Println(r)
// }

// func isPalindrome(s string) bool {

// 	nonAlphanumericRegex := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

// 	s = nonAlphanumericRegex.ReplaceAllString(s, "")
// 	s = strings.ToLower(s)
// 	s = strings.ReplaceAll(s, " ", "")
// 	// s = strings.ReplaceAll(s, ",", "")
// 	// s = strings.ReplaceAll(s, ":", "")
// 	// fmt.Println(s)
// 	if len(s) == 0 {
// 		return true
// 	}

// 	right := len(s) - 1
// 	for i := 0; i < len(s)/2; i++ {
// 		if s[i] == s[right] {

// 			right--
// 		} else {
// 			return false
// 		}
// 	}

// 	return true
// }
