package main

import "fmt"

// LCR 146. 螺旋遍历二维数组
//给定一个二维数组 array，请返回「螺旋遍历」该数组的结果。
//螺旋遍历：从左上角开始，按照 向右、向下、向左、向上 的顺序 依次 提取元素，然后再进入内部一层重复相同的步骤，
//直到提取完所有元素。
//
//示例 1：
//输入：array = [[1,2,3],[8,9,4],[7,6,5]]
//输出：[1,2,3,4,5,6,7,8,9]
//
//示例 2：
//输入：array  = [[1,2,3,4],[12,13,14,5],[11,16,15,6],[10,9,8,7]]
//输出：[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16]
//
//限制：
//0 <= array.length <= 100
//0 <= array[i].length <= 100

func main() {
	ans := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	ans = [][]int{[]int{6, 7, 9}}

	spiralArray(ans)
	var n int
	for {
		if _, err := fmt.Scan(&n); err != nil {
			break
		}
		ans := make([][]int, n)
		for i := 0; i < n; i++ {
			ans[i] = make([]int, n)
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				fmt.Scan(&ans[i][j])
			}
		}
		spiralArray(ans)
	}
}

// 错误点： 可能不是正方形
func spiralArray(array [][]int) []int {
	startX, startY := 0, 0
	lenX := len(array)
	lenY := 0
	if lenX > 0 {
		lenY = len(array[0])
	}
	loop, circleCount := lenX/2, 1

	res := make([]int, lenX*lenY)
	index := 0

	for loop >= 0 {
		// 上横
		for j := startY; j < lenY-circleCount && j < lenY && index < len(res); j++ {
			res[index] = array[startX][j]
			index++
		}
		// 右竖
		for i := startX; i < lenX-circleCount && i < lenX && index < len(res); i++ {
			res[index] = array[i][lenY-circleCount]
			index++
		}
		// 下横
		for j := lenY - circleCount; j > startY && j < lenY && index < len(res); j-- {
			res[index] = array[lenX-circleCount][j]
			index++
		}
		// 左竖
		for i := lenX - circleCount; i > startX && i < lenX && index < len(res); i-- {
			res[index] = array[i][startY]
			index++
		}

		startX++
		startY++
		loop--
		circleCount++
	}
	if lenX == lenY && lenX%2 == 1 {
		res[len(res)-1] = array[lenX/2][lenY/2]
	}
	return res
}

func spiralArray_ans(array [][]int) []int {
	if len(array) == 0 || len(array[0]) == 0 {
		return []int{}
	}
	rows, columns := len(array), len(array[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, columns)
	}

	var (
		total          = rows * columns
		order          = make([]int, total)
		row, column    = 0, 0
		directions     = [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
		directionIndex = 0
	)

	for i := 0; i < total; i++ {
		order[i] = array[row][column]
		visited[row][column] = true
		nextRow, nextColumn := row+directions[directionIndex][0], column+directions[directionIndex][1]
		if nextRow < 0 || nextRow >= rows || nextColumn < 0 || nextColumn >= columns || visited[nextRow][nextColumn] {
			directionIndex = (directionIndex + 1) % 4
		}
		row += directions[directionIndex][0]
		column += directions[directionIndex][1]
	}
	return order
}
