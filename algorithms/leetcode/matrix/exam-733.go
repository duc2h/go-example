package main

import "fmt"

// https://leetcode.com/problems/flood-fill/

// An image is represented by an m x n integer grid image where image[i][j] represents
// the pixel value of the image.
// You are also given three integers sr, sc, and color. You should perform a flood fill on
// the image starting from the pixel image[sr][sc].
// To perform a flood fill, consider the starting pixel, plus any pixels connected 4-directionally
// to the starting pixel of the same color as the starting pixel, plus any pixels connected 4-directionally
// to those pixels (also with the same color), and so on. Replace the color of all of the aforementioned
// pixels with color.
// Return the modified image after performing the flood fill.

// Input: image = [[1,1,1],[1,1,0],[1,0,1]], sr = 1, sc = 1, color = 2
// Output: [[2,2,2],[2,2,0],[2,0,1]]
// Explanation: From the center of the image with position (sr, sc) = (1, 1) (i.e., the red pixel), all pixels connected by a path of the same color as the starting pixel (i.e., the blue pixels) are colored with the new color.
// Note the bottom corner is not colored 2, because it is not 4-directionally connected to the starting pixel.

// Example 2:
// Input: image = [[0,0,0],[0,0,0]], sr = 0, sc = 0, color = 0
// Output: [[0,0,0],[0,0,0]]
// Explanation: The starting pixel is already colored 0, so no changes are made to the image.

func main() {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}

	r := floodFill(image, 2, 2, 2)
	fmt.Println(r)
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	sourceColor := image[sr][sc]
	if sourceColor != color {
		dfs(image, sr, sc, sourceColor, color)
	}

	return image
}

func dfs(image [][]int, x, y, sourceColor, newColor int) {
	if x < 0 || x >= len(image) || y < 0 || y >= len(image[0]) || image[x][y] != sourceColor {
		return
	}
	image[x][y] = newColor
	dfs(image, x+1, y, sourceColor, newColor)
	dfs(image, x-1, y, sourceColor, newColor)
	dfs(image, x, y+1, sourceColor, newColor)
	dfs(image, x, y-1, sourceColor, newColor)

}
