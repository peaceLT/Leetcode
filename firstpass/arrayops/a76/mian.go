package main

import (
	"fmt"
	"math"
)

//76. 最小覆盖子串
//给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
//注意：
//对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
//如果 s 中存在这样的子串，我们保证它是唯一的答案。
//
//
//示例 1：
//输入：s = "ADOBECODEBANC", t = "ABC"
//输出："BANC"
//解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
//
//示例 2：
//输入：s = "a", t = "a"
//输出："a"
//解释：整个字符串 s 是最小覆盖子串。
//
//示例 3:
//输入: s = "a", t = "aa"
//输出: ""
//解释: t 中两个字符 'a' 均应包含在 s 的子串中，
//因此没有符合条件的子字符串，返回空字符串。
//
//
//提示：
//m == s.length
//n == t.length
//1 <= m, n <= 105
//s 和 t 由英文字母组成

func main() {
	var s, t string
	for {
		if _, err := fmt.Scan(&s, &t); err != nil {
			break
		}
		fmt.Println(minWindow(s, t))
	}
}

func minWindow_mine(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	distance := 0
	subStr := make([]uint8, 0, len(t))
	minstr := ""
	left, right := 0, 0
	tCount, sCount := map[int32]int{}, map[int32]int{}
	for _, v := range t {
		tCount[v]++
	}

	for left < len(s) && right < len(s) {
		subStr = append(subStr, s[right])
		if sCount[int32(s[right])] < tCount[int32(s[right])] {
			distance++
		}
		if _, ok := tCount[int32(s[right])]; ok {
			sCount[int32(s[right])]++
		}

		// 找到
		for distance == len(t) {
			if len(minstr) > len(subStr) || len(minstr) == 0 {
				minstr = string(subStr)
			}

			if _, ok := tCount[int32(s[left])]; ok && sCount[int32(s[left])] == tCount[int32(s[left])] {
				distance--
				sCount[int32(s[left])]--
			} else if sCount[int32(s[left])] > tCount[int32(s[left])] {
				sCount[int32(s[left])]--
			}
			left++
			subStr = subStr[1:]
		}
		right++
	}

	return minstr
}

// minWindow 寻找包含字符串t的全部字符的最短子串。
// 参数s: 源字符串
// 参数t: 目标字符串，包含需要查找的字符
// 返回值: 最短子串
func minWindow(s string, t string) string {
	// 初始化两个字典，ori存储目标字符串t中每个字符的出现次数，
	// cnt用于在遍历过程中存储当前窗口内每个字符的出现次数
	ori, cnt := map[byte]int{}, map[byte]int{}

	// 统计目标字符串t中每个字符的出现次数
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	// sLen用于存储源字符串s的长度
	sLen := len(s)
	// 初始化最短子串的长度为最大整数，以便后续寻找更短的子串
	len := math.MaxInt32
	// ansL和ansR用于记录最短子串的左右索引
	ansL, ansR := -1, -1

	// check函数用于检查当前窗口是否包含了目标字符串t的所有字符
	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}

	// 使用双指针l和r来滑动窗口，寻找最短子串
	for l, r := 0, 0; r < sLen; r++ {
		// 如果当前字符是目标字符串t中的字符，则增加cnt中对应字符的计数
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}

		// 当当前窗口包含了目标字符串t的所有字符时，尝试缩小窗口
		for check() && l <= r {
			// 如果当前窗口的长度小于已记录的最短长度，则更新最短长度和索引
			if r-l+1 < len {
				len = r - l + 1
				ansL, ansR = l, l+len
			}
			// 移动左指针l，尝试缩小窗口，同时更新cnt中对应字符的计数
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}

	// 如果没有找到合适的子串，则返回空字符串
	if ansL == -1 {
		return ""
	}
	// 根据最短子串的索引返回子串
	return s[ansL:ansR]
}

func minWindow2(s string, t string) string {
	if len(t) == 0 {
		return ""
	}

	// 初始化 ori 和 cnt
	ori := make(map[byte]int)
	cnt := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	// 初始化变量
	sLen := len(s)
	minLen := math.MaxInt32
	ansL, ansR := -1, -1
	matchCount := 0 // 记录匹配的字符数量

	// 主循环
	l := 0
	for r := 0; r < sLen; r++ {
		// 更新右边界
		if _, ok := ori[s[r]]; ok {
			cnt[s[r]]++
			if cnt[s[r]] == ori[s[r]] {
				matchCount++
			}
		}

		// 尝试收缩左边界
		for matchCount == len(ori) && l <= r {
			if r-l+1 < minLen {
				minLen = r - l + 1
				ansL, ansR = l, l+minLen
			}

			// 更新左边界
			if _, ok := ori[s[l]]; ok {
				if cnt[s[l]] == ori[s[l]] {
					matchCount--
				}
				cnt[s[l]]--
			}
			l++
		}
	}

	// 返回结果
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}
