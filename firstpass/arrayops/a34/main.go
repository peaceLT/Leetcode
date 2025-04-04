package main

import (
	"fmt"
	"sort"
)

//在排序数组中查找元素的第一个和最后一个位置
//给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
//
//如果数组中不存在目标值 target，返回 [-1, -1]。
//
//进阶：你可以设计并实现时间复杂度为 $O(\log n)$ 的算法解决此问题吗？
//
//示例 1：
//
//输入：nums = [5,7,7,8,8,10], target = 8
//输出：[3,4]
//示例 2：
//
//输入：nums = [5,7,7,8,8,10], target = 6
//输出：[-1,-1]
//示例 3：
//
//输入：nums = [], target = 0
//输出：[-1,-1]

func main() {
	var nums []int
	for {
		num := 0
		if _, err := fmt.Scan(&num); err != nil {
			break
		}
		nums = append(nums, num)
	}
	target := nums[len(nums)-1]
	nums = nums[:len(nums)-1]
	res := searchRange(nums, target)
	fmt.Println(res)
}

// O(Logn)
func searchRange_Mine(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1

	index := -1
	for left <= right {
		mid := left + (right-left)>>1

		if nums[mid] == target {
			index = mid
			break
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	res := make([]int, 0, 2)
	if index != -1 {
		res = findMinMax(index, nums)
	} else {
		res = append(res, -1, -1)
	}

	return res
}

func findMinMax(index int, nums []int) []int {
	target := nums[index]
	res := []int{len(nums), 0}
	for i := index; i >= 0; i-- {
		if nums[i] == target && i <= res[0] {
			res[0] = i
		} else {
			break
		}
	}
	for i := index; i < len(nums); i++ {
		if nums[i] == target && i >= res[1] {
			res[1] = i
		} else {
			break
		}
	}
	return res
}

// 参考答案
func searchRange_kai(nums []int, target int) []int {
	leftBorder := getLeft(nums, target)
	rightBorder := getRight(nums, target)
	// 情况一
	if leftBorder == -2 || rightBorder == -2 {
		return []int{-1, -1}
	}
	// 情况三
	if rightBorder-leftBorder > 1 {
		return []int{leftBorder + 1, rightBorder - 1}
	}
	// 情况二
	return []int{-1, -1}
}

func getLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	border := -2        // 记录border没有被赋值的情况；这里不能赋值-1，target = num[0]时，会无法区分情况一和情况二
	for left <= right { // []闭区间
		mid := left + ((right - left) >> 1)
		if nums[mid] >= target { // 找到第一个等于target的位置
			right = mid - 1
			border = right
		} else {
			left = mid + 1
		}
	}
	return border
}

func getRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	border := -2
	for left <= right {
		mid := left + ((right - left) >> 1)
		if nums[mid] > target {
			right = mid - 1
		} else { // 找到第一个大于target的位置
			left = mid + 1
			border = left
		}
	}
	return border
}

// 官方题解
func searchRange(nums []int, target int) []int {
	leftmost := sort.SearchInts(nums, target)
	// 没找到
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return []int{leftmost, rightmost}
}
