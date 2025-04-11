package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// 44. 开发商购买土地（第五期模拟笔试）
//题目描述
//在一个城市区域内，被划分成了n * m个连续的区块，每个区块都拥有不同的权值，代表着其土地价值。目前，有两家开发公司，A 公司和 B 公司，
//希望购买这个城市区域的土地。
//
//现在，需要将这个城市区域的所有区块分配给 A 公司和 B 公司。
//
//然而，由于城市规划的限制，只允许将区域按横向或纵向划分成两个子区域，而且每个子区域都必须包含一个或多个区块。
//为了确保公平竞争，你需要找到一种分配方式，使得 A 公司和 B 公司各自的子区域内的土地总价值之差最小。
//
//注意：区块不可再分。
//
//输入描述
//第一行输入两个正整数，代表 n 和 m。
//
//接下来的 n 行，每行输出 m 个正整数。
//
//输出描述
//请输出一个整数，代表两个子区域内土地总价值之间的最小差距。
//输入示例
//3 3
//1 2 3
//2 1 3
//1 2 3
//输出示例
//0
//提示信息
//如果将区域按照如下方式划分：
//
//1 2 | 3
//2 1 | 3
//1 2 | 3
//
//两个子区域内土地总价值之间的最小差距可以达到 0。
//
//数据范围：
//
//1 <= n, m <= 100；
//n 和 m 不同时为 1。

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	line := scanner.Text()
	m, n := 0, 0
	_, err := fmt.Sscanf(line, "%d %d", &n, &m)
	if err != nil {
		return
	}
	land := make([][]int, n)
	for i := 0; i < n; i++ {
		land[i] = make([]int, m)
		scanner.Scan()
		line = scanner.Text()
		line = strings.TrimSpace(line)
		values := strings.Split(line, " ")
		for j := 0; j < m; j++ {
			value, _ := strconv.Atoi(values[j])
			land[i][j] = value
		}
	}
	fmt.Println(land)

	preMaxtrix := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		preMaxtrix[i] = make([]int, m+1)
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			preMaxtrix[i][j] = land[i-1][j-1] + preMaxtrix[i-1][j] + preMaxtrix[i][j-1] - preMaxtrix[i-1][j-1]
		}
	}

	totalSum := preMaxtrix[n][m]

	minDiff := math.MaxInt32

	for i := 1; i < n; i++ {
		topSum := preMaxtrix[i][m]
		bottomSum := totalSum - topSum

		diff := int(math.Abs(float64(topSum - bottomSum)))
		if diff < minDiff {
			minDiff = diff
		}
	}
	for j := 1; j < m; j++ {
		topSum := preMaxtrix[n][j]
		bottomSum := totalSum - topSum

		diff := int(math.Abs(float64(topSum - bottomSum)))
		if diff < minDiff {
			minDiff = diff
		}
	}

	fmt.Println(minDiff)
}
