package main

import "fmt"

// 904. 水果成篮
//你正在探访一家农场，农场从左到右种植了一排果树。这些树用一个整数数组 fruits 表示，其中 fruits[i] 是第 i 棵树上的水果 种类 。
//
//你想要尽可能多地收集水果。然而，农场的主人设定了一些严格的规矩，你必须按照要求采摘水果：
//
//你只有 两个 篮子，并且每个篮子只能装 单一类型 的水果。每个篮子能够装的水果总量没有限制。
//你可以选择任意一棵树开始采摘，你必须从 每棵 树（包括开始采摘的树）上 恰好摘一个水果 。采摘的水果应当符合篮子中的水果类型。每采摘一次，你将会向右移动到下一棵树，并继续采摘。
//一旦你走到某棵树前，但水果不符合篮子的水果类型，那么就必须停止采摘。
//给你一个整数数组 fruits ，返回你可以收集的水果的 最大 数目。
//
//
//
//示例 1：
//
//输入：fruits = [1,2,1]
//输出：3
//解释：可以采摘全部 3 棵树。
//示例 2：
//
//输入：fruits = [0,1,2,2]
//输出：3
//解释：可以采摘 [1,2,2] 这三棵树。
//如果从第一棵树开始采摘，则只能采摘 [0,1] 这两棵树。
//示例 3：
//
//输入：fruits = [1,2,3,2,2]
//输出：4
//解释：可以采摘 [2,3,2,2] 这四棵树。
//如果从第一棵树开始采摘，则只能采摘 [1,2] 这两棵树。
//示例 4：
//
//输入：fruits = [3,3,3,1,2,1,1,2,3,3,4]
//输出：5
//解释：可以采摘 [1,2,1,1,2] 这五棵树。
//
//
//提示：
//
//1 <= fruits.length <= 10^5
//0 <= fruits[i] < fruits.length

func main() {
	fmt.Println(totalFruit([]int{1, 2, 1}))
	fmt.Println(totalFruit([]int{0, 1, 2, 2}))
	fmt.Println(totalFruit([]int{1, 2, 3, 2, 2}))
}

func totalFruit_error(fruits []int) int {
	maxNum, sum, i, fstType, sedType := 0, 0, 0, -1, -1
	isbool := false

	for j := 0; j < len(fruits); j++ {
		if fstType == -1 || fruits[j] == fstType {
			sum++
			fstType = fruits[j]
		} else if sedType == -1 || fruits[j] == sedType {
			sum++
			sedType = fruits[j]
		} else {
			if sum > maxNum {
				maxNum = sum
			}

			if isbool {
				sum, i = 1, j
				isbool = false
				fstType, sedType = fruits[j], -1
			} else {
				for fruits[i] == fstType {
					sum--
					i++
				}
				fstType, sedType = fruits[j], -1
			}
		}
	}

	return maxNum
}

// 我们可以使用滑动窗口解决本题，left 和 right 分别表示满足要求的窗口的左右边界，同时我们使用哈希表存储这个窗口内的数以及出现的次数。
//
// 我们每次将 right 移动一个位置，并将 fruits[right] 加入哈希表。如果此时哈希表不满足要求（即哈希表中出现超过两个键值对），那么我们需要不断移动 left，并将 fruits[left] 从哈希表中移除，直到哈希表满足要求为止。
//
// 需要注意的是，将 fruits[left] 从哈希表中移除后，如果 fruits[left] 在哈希表中的出现次数减少为 0，需要将对应的键值对从哈希表中移除
func totalFruit(fruits []int) int {
	cnt := map[int]int{}
	left, ans := 0, 0
	for right, x := range fruits {
		cnt[x]++
		for len(cnt) > 2 {
			y := fruits[left]
			cnt[y]--
			if cnt[y] == 0 {
				delete(cnt, y)
			}
			left++
		}
		if right-left+1 > ans {
			ans = right - left + 1
		}
	}
	return ans
}
