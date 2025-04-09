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

func minWindow(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	len := math.MaxInt32
	ansL, ansR := -1, -1

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < len {
				len = r - l + 1
				ansL, ansR = l, l+len
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}
