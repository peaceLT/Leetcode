package main

import (
	. "fmt"
)

//二分查找
//给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
//
//示例 1:
//输入: nums = [-1,0,3,5,9,12], target = 9
//输出: 4
//解释: 9 出现在 nums 中并且下标为 4
//
//示例 2:
//输入: nums = [-1,0,3,5,9,12], target = 2
//输出: -1
//解释: 2 不存在 nums 中因此返回 -1

func main() {
	var nums []int
	var target int
	for {
		var num int
		if _, err := Scan(&num); err != nil {
			break
		}
		nums = append(nums, num)
	}
	target = nums[len(nums)-1]
	nums = nums[:len(nums)-1]
	resIndex := search(nums, target)
	Println(resIndex)
}

func search(nums []int, target int) int {
	// 初始化左右边界
	left := 0
	right := len(nums) - 1

	// 循环逐步缩小区间范围
	for left <= right {
		mid := left + (right-left)>>1

		// 根据 nums[mid] 和 target 的大小关系
		// 调整区间范围
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// 未找到
	return -1
}
