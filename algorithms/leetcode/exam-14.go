package main

// https://leetcode.com/problems/longest-common-prefix/

// func main() {
// 	s := []string{"flight", "flower", "flow", "flight"}
// 	// s1 := []string{"dog", "racecar", "car"}
// 	// s2 := []string{"", "b"}
// 	// s3 := []string{"cir", "car"}

// 	fmt.Println(longestCommonPrefix(s))
// }

func longestCommonPrefix(strs []string) string {
	p := strs[0]
	for _, s := range strs {
		i := 0
		for j := 0; j < len(s); j++ {
			if len(s) > j && len(p) > j && s[j] == p[j] {
				i++
			} else {
				break
			}
		}

		p = p[:i]
	}

	return p
}
