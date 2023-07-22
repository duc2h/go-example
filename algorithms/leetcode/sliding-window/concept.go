package main

// https://www.scaler.com/topics/sliding-window-algorithm/

// We have n = 5, arr[5] = {10,20,10,30,5} and m = 3 Then the result would be 60 (20+10+30).
// We have n = 7, arr[7] = {2,10,17,1,9,13,4} and m = 4 Then the result would be 40 (17+1+9+13).

// func main() {
// 	a := []int{2, 10, 17, 1, 9, 13, 4}
// 	r := test(a, 4)

// 	fmt.Println(r)
// }

func test(a []int, k int) int {
	var (
		length = len(a)
		max    = 0
	)

	for i := 0; i < k; i++ {
		max += a[i]
	}

	for i := 0; i < length; i++ {
		x := i + k
		if x >= length {
			break
		}
		temp := max + a[x] - a[i]
		max = findMax(temp, max)
	}

	return max
}

func maxScore1(cardPoints []int, k int) int {
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
