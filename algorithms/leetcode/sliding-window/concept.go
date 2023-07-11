package main

import "fmt"

// https://www.scaler.com/topics/sliding-window-algorithm/

// We have n = 5, arr[5] = {10,20,10,30,5} and m = 3 Then the result would be 60 (20+10+30).
// We have n = 7, arr[7] = {2,10,17,1,9,13,4} and m = 4 Then the result would be 40 (17+1+9+13).

func main() {
	a := []int{10, 20, 10, 30, 5}
	r := maxScore(a, 3)

	fmt.Println(r)
}

func maxScore(cardPoints []int, k int) int {
	var (
		length = len(cardPoints)
		max    = 0
	)

	for i := 0; i < k; i++ {
		max += cardPoints[i]
	}

	for i := k; i < length; i++ {
		temp := max + cardPoints[i]
		temp = temp - cardPoints[i-k]

		max = findMax(max, temp)
	}

	return max
}

func findMax(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}

	return num2
}
