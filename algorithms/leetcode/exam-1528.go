package main

// https://leetcode.com/problems/shuffle-string/

// func main() {
// 	s := "codeleet"
// 	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}
// 	fmt.Println(restoreString(s, indices))
// }

func restoreString(s string, indices []int) string {
	result := make([]byte, len(s))
	for i, indice := range indices {
		result[indice] = s[i]
	}

	return string(result)

}
