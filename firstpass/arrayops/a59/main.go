package main

import "fmt"

//59. 螺旋矩阵 II
//给你一个正整数 n ，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
//
//
//示例 1：
//输入：n = 3
//输出：[[1,2,3],[8,9,4],[7,6,5]]

//示例 2：
//输入：n = 1
//输出：[[1]]
//
//
//提示：
//1 <= n <= 20

func main() {
	var n int
	for {
		if _, err := fmt.Scan(&n); err != nil {
			break
		}
		generateMatrix(n)
	}
}

func generateMatrix(n int) [][]int {
	startX, startY := 0, 0
	loop := n / 2
	num, circulCount := 1, 1
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	for loop > 0 {
		// 上横
		for j := startY; j < n-circulCount; j++ {
			res[startX][j] = num
			num++
		}
		// 右竖
		for i := startX; i < n-circulCount; i++ {
			res[i][n-circulCount] = num
			num++
		}
		// 下横
		for j := n - circulCount; j > startY; j-- {
			res[n-circulCount][j] = num
			num++
		}
		// 左竖
		for i := n - circulCount; i > startX; i-- {
			res[i][startY] = num
			num++
		}

		circulCount++
		startX++
		startY++
		loop--
	}
	// 单数
	if n%2 == 1 {
		res[n/2][n/2] = n * n
	}
	return res
}
