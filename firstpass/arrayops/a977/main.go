package main

import (
	"sort"
)

//给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
//
//
//
//示例 1：
//
//输入：nums = [-4,-1,0,3,10]
//输出：[0,1,9,16,100]
//解释：平方后，数组变为 [16,1,0,9,100]
//排序后，数组变为 [0,1,9,16,100]
//示例 2：
//
//输入：nums = [-7,-3,2,3,11]
//输出：[4,9,9,49,121]
//
//
//提示：
//
//1 <= nums.length <= 10^4
//-10^4 <= nums[i] <= 10^4
//nums 已按 非递减顺序 排序
//
//
//进阶：
//请你设计时间复杂度为 O(n) 的算法解决本问题

func sortedSquares(nums []int) []int {
	left, right := 0, len(nums)-1
	ans := make([]int, len(nums), len(nums))

	for i := len(nums) - 1; left <= right; i-- {
		if nums[left]*nums[left] >= nums[right]*nums[right] {
			ans[i] = nums[left] * nums[left]
			left++
		} else {
			ans[i] = nums[right] * nums[right]
			right--
		}
	}
	return ans
}

// 官方答案——双指针
func sortedSquares_official1(nums []int) []int {
	n := len(nums)
	lastNegIndex := -1
	for i := 0; i < n && nums[i] < 0; i++ {
		lastNegIndex = i
	}

	ans := make([]int, 0, n)
	for i, j := lastNegIndex, lastNegIndex+1; i >= 0 || j < n; {
		if i < 0 {
			ans = append(ans, nums[j]*nums[j])
			j++
		} else if j == n {
			ans = append(ans, nums[i]*nums[i])
			i--
		} else if nums[i]*nums[i] < nums[j]*nums[j] {
			ans = append(ans, nums[i]*nums[i])
			i--
		} else {
			ans = append(ans, nums[j]*nums[j])
			j++
		}
	}

	return ans
}

// 暴力
func sortedSquares_official2(nums []int) []int {
	ans := make([]int, len(nums))
	for i, v := range nums {
		ans[i] = v * v
	}
	sort.Ints(ans)
	return ans
}
