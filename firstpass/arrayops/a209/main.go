package main

import (
	"fmt"
	"sort"
)

//209.长度最小的子数组
//
//给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。
//
//示例：
//
//输入：s = 7, nums = [2,3,1,2,4,3]
//输出：2
//解释：子数组 [4,3] 是该条件下的长度最小的子数组。
//提示：
//
//1 <= target <= 10^9
//1 <= nums.length <= 10^5
//1 <= nums[i] <= 10^5

func main() {
	//for {
	//	var n, s int
	//	if _, err := fmt.Scan(&n, &s); err != nil {
	//		break
	//	}
	//
	//	nums := make([]int, n, n)
	//	for i := 0; i < n; i++ {
	//		fmt.Scan(&nums[i])
	//	}
	//	fmt.Println(minSubArrayLen(s, nums))
	//}

	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
}

// 错误思路： 找出该数组中满足其总和大于等于 target 的长度最小的子数组 [numsl, numsl+1, ..., numsr-1, numsr] ==> 由此可见，需要连续数字组成的数组，以下的方法改变了原数组的顺序
func minSubArrayLen_error(target int, nums []int) int {
	sort.Ints(nums)
	sum, cnt := 0, 0

	for i := len(nums) - 1; i >= 0; i-- {
		if sum < target {
			sum += nums[i]
			cnt++
		}
		if sum >= target {
			return cnt
		}
	}
	return 0
}

func minSubArrayLen(target int, nums []int) int {
	minLen, subLen, sum, index := len(nums)+1, 0, 0, 0

	// 寻找窗口结束
	for i := index; i < len(nums); i++ {
		sum += nums[i]
		subLen++

		if sum >= target {
			if subLen < minLen {
				minLen = subLen
			}
			sum, subLen = 0, 0
			index++
			i = index
		}
	}

	if minLen > len(nums) {
		return 0
	}
	return minLen
}

func minSubArrayLen_ka(target int, nums []int) int {
	i := 0
	l := len(nums)  // 数组长度
	sum := 0        // 子数组之和
	result := l + 1 // 初始化返回长度为l+1，目的是为了判断“不存在符合条件的子数组，返回0”的情况
	for j := 0; j < l; j++ {
		sum += nums[j]
		for sum >= target {
			subLength := j - i + 1
			if subLength < result {
				result = subLength
			}
			sum -= nums[i]
			i++
		}
	}
	if result == l+1 {
		return 0
	} else {
		return result
	}
}

func minSubArrayLen_ka2(target int, nums []int) int {
	i, sum, minLen, subLen := 0, 0, len(nums), 0
	for j := 0; j < len(nums); j++ {
		sum += nums[j]

		for sum >= target {
			subLen = j - i + 1
			if subLen < minLen {
				minLen = subLen
			}
			sum -= nums[i]
			i--
		}
	}

	if minLen == len(nums)+1 {
		return 0
	}
	return minLen
}
