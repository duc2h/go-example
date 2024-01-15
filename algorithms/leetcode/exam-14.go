package main

// https://leetcode.com/problems/longest-common-prefix/

// func main() {
// 	// s := []string{"flight", "flower", "flow", "flight"}
// 	// s1 := []string{"dog", "racecar", "car"}
// 	// s2 := []string{"", "b"}
// 	s3 := []string{"cir", "car"}

// 	fmt.Println(longestCommonPrefix(s3))
// }

// func longestCommonPrefix(strs []string) string {
// 	p := strs[0]
// 	for _, s := range strs {
// 		i := 0
// 		for j := 0; j < len(s); j++ {
// 			if len(s) > j && len(p) > j && s[j] == p[j] {
// 				i++
// 			} else {
// 				break
// 			}
// 		}

// 		p = p[:i]
// 	}

// 	return p
// }

func longestCommonPrefix(strs []string) string {
	result := strs[0]

	for _, str := range strs {
		temp := ""
		for i := 0; i < len(result); i++ {
			if i < len(str) {
				s1 := result[i]
				s2 := str[i]

				if s1 == s2 {
					temp = temp + string(s1)
				} else {
					if len(result) > len(temp) {
						result = temp
					}
				}
			} else {
				result = str
			}
		}
	}

	return result
}
