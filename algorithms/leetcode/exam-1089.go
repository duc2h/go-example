package main

// https://leetcode.com/problems/duplicate-zeros/

// func main() {
// 	arr := []int{1}
// 	duplicateZeros(arr)

// 	fmt.Println(arr)
// }

func duplicateZeros(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == 0 {

			for j := len(arr) - 1; j > i+1; j-- {
				arr[j] = arr[j-1]
			}

			i++
			if i < len(arr) {

				arr[i] = 0
			}
		}
	}

}
