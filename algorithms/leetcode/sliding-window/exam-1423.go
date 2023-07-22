package main

import "fmt"

// https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/

// There are several cards arranged in a row, and each card has an associated number of points. The points are given in the integer array cardPoints.
// In one step, you can take one card from the beginning or from the end of the row. You have to take exactly k cards.
// Your score is the sum of the points of the cards you have taken.
// Given the integer array cardPoints and the integer k, return the maximum score you can obtain.

// Example 1:
// Input: cardPoints = [1,2,3,4,5,6,1], k = 3
// Output: 12
// Explanation: After the first step, your score will always be 1. However, choosing the rightmost card first will maximize your total score. The optimal strategy is to take the three cards on the right, giving a final score of 1 + 6 + 5 = 12.

// Example 2:
// Input: cardPoints = [2,2,2], k = 2
// Output: 4
// Explanation: Regardless of which two cards you take, your score will always be 4.

// Example 3:
// Input: cardPoints = [9,7,7,9,7,7,9], k = 7
// Output: 55
// Explanation: You have to take all the cards. Your score is the sum of points of all cards.

// Example 4:
// Input: [100, 40, 17, 9, 73, 75] k=3
// Output: 248 = 100 + 73 + 75.
// 100 + 40 + 17
// 100 + 40 + 75
// 100 + 75+ 73 => biggest
// 75 + 73 + 9

// it is related to this question: https://leetcode.com/discuss/interview-question/701938/google-max-sum-of-k-elements

// func main() {
// 	a := []int{100, 40, 17, 9, 73, 75}
// 	r := maxScore(a, 3)
// 	fmt.Println(r)
// }

// explanation: 1, 2, 3, 4, 5, 6, 1
// step1: max = 6, left = 2, right = 6.
// step2: i =0, temp = 4, max = 6, left = 1, right = 5.
// step3: i=1, temp = 8, max = 8, left = 0, right = 4.
// step4: i = 2, temp = 12, max = 12, left = -1, right = 3. done.
func maxScore(cardPoints []int, k int) int {
	length := len(cardPoints)
	max := 0
	left, right := k-1, length-1

	fmt.Println(left)

	for i := 0; i < k; i++ {
		max += cardPoints[i]
	}
	temp := max

	for i := 0; i < k; i++ {
		numL := cardPoints[left]
		numR := cardPoints[right]

		temp = temp + numR - numL
		max = finMax1(temp, max)

		left--
		right--
	}

	return max
}

func finMax1(num1, num2 int) int {

	if num1 > num2 {
		return num1
	}

	return num2
}
