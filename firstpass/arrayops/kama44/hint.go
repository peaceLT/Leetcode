package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// main_ka是程序的入口点。
func main_ka() {
	// n和m分别表示土地的行数和列数。
	var n, m int

	// 使用bufio.NewReader读取标准输入。
	reader := bufio.NewReader(os.Stdin)

	// 读取并解析输入的第一行，获取土地的行数和列数。
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	params := strings.Split(line, " ")
	n, _ = strconv.Atoi(params[0])
	m, _ = strconv.Atoi(params[1]) //n和m读取完成

	// 初始化表示土地的二维切片。
	land := make([][]int, n) //土地矩阵初始化

	// 逐行读取并解析土地的每个单元格的值。
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		values := strings.Split(line, " ")
		land[i] = make([]int, m)
		for j := 0; j < m; j++ {
			value, _ := strconv.Atoi(values[j])
			land[i][j] = value
		}
	} //所有读取完成

	// 初始化用于计算前缀和的二维切片。
	preMatrix := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		preMatrix[i] = make([]int, m+1)
	}

	// 计算土地的前缀和矩阵。
	for a := 1; a < n+1; a++ {
		for b := 1; b < m+1; b++ {
			preMatrix[a][b] = land[a-1][b-1] + preMatrix[a-1][b] + preMatrix[a][b-1] - preMatrix[a-1][b-1]
		}
	}

	// totalSum是土地上所有单元格值的总和。
	totalSum := preMatrix[n][m]

	// 初始化minDiff为最大整数，以便后续寻找最小差值。
	minDiff := math.MaxInt32 //初始化极大数，用于比较

	// 遍历每一行，计算将土地按行分割成两部分的差值，并更新最小差值。
	for i := 1; i < n; i++ {
		topSum := preMatrix[i][m]

		bottomSum := totalSum - topSum

		diff := int(math.Abs(float64(topSum - bottomSum)))
		if diff < minDiff {
			minDiff = diff
		}
	}

	// 遍历每一列，计算将土地按列分割成两部分的差值，并更新最小差值。
	for j := 1; j < m; j++ {
		topSum := preMatrix[n][j]

		bottomSum := totalSum - topSum

		diff := int(math.Abs(float64(topSum - bottomSum)))
		if diff < minDiff {
			minDiff = diff
		}
	}

	// 输出最小差值。
	fmt.Println(minDiff)
}
