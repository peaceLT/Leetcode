package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//58. 区间和（第九期模拟笔试）
//题目描述
//给定一个整数数组 Array，请计算该数组在每个指定区间内元素的总和。
//输入描述
//第一行输入为整数数组 Array 的长度 n，接下来 n 行，每行一个整数，表示数组的元素。随后的输入为需要计算总和的区间下标：a，b （b > = a），直至文件结束。
//输出描述
//输出每个指定区间内元素的总和。
//输入示例
//5
//1
//2
//3
//4
//5
//0 1
//1 3
//输出示例
//3
//9
//提示信息
//数据范围：
//0 < n <= 100000

func main_mine() {
	var n int
	fmt.Scan(&n)
	ans, pre := make([]int, n), make([]int, n)
	fmt.Scan(&ans[0])
	pre[0] = ans[0]
	for i := 1; i < n; i++ {
		fmt.Scan(&ans[i])
		pre[i] = pre[i-1] + ans[i]
	}
	for {
		var left, right int
		if _, err := fmt.Scan(&left, &right); err != nil {
			break
		}
		if left == 0 {
			fmt.Println(pre[right])
		} else if left > 0 {
			fmt.Println(pre[right] - pre[left-1])
		}
	}
}

func calculateRangeSum(array []int, start, end int) int {
	prefixSumStart := 0
	if start > 0 {
		prefixSumStart = array[start-1]
	}
	return array[end] - prefixSumStart
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// 读取数组大小
	scanner.Scan()
	size, _ := strconv.Atoi(scanner.Text())

	// 读取数组元素
	numberArray := make([]int, size)
	for i := 0; i < size; i++ {
		scanner.Scan()
		numberArray[i], _ = strconv.Atoi(scanner.Text())
	}

	// 计算前缀和数组
	prefixArray := make([]int, size)
	prefixArray[0] = numberArray[0]
	for i := 1; i < size; i++ {
		prefixArray[i] = prefixArray[i-1] + numberArray[i]
	}

	// 计算区间和
	for scanner.Scan() {
		line := scanner.Text()
		fields := []int{0, 0}
		_, err := fmt.Sscanf(line, "%d %d", &fields[0], &fields[1])
		if err != nil {
			continue
		}

		if fields[0] < 0 || fields[1] >= size {
			continue
		}
		fmt.Println(calculateRangeSum(prefixArray, fields[0], fields[1]))
		// 		fmt.Println(calculateRangeSum2(numberArray, fields[0], fields[1]))
	}
}
