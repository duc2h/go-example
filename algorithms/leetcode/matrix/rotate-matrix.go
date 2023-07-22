package main

import (
	"fmt"
)

// https://www.enjoyalgorithms.com/blog/rotate-a-matrix-by-90-degrees-in-an-anticlockwise-direction

// Given a matrix n x n. Rotate the matrix 90 degrees to the left.

// Example 1:
// input
// 1 2 3
// 4 5 6
// 7 8 9

// output:
// 3 6 9
// 2 5 8
// 1 4 7

// Example 2:
// input:
// 1 2 3 4
// 5 6 7 8
// 9 10 11 12
// 13 14 15 16

// output:
// 4 8 12 16
// 3 7 11 15
// 2 6 10 14
// 1 5 9 13

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	r := rotateMatrix(matrix, 4)

	fmt.Println(r)
}

func rotateMatrix(matrix [][]int, n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result[n-j-1][i] = matrix[i][j]
		}
	}

	return result
}
