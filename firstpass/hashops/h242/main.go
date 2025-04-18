package main

import (
	"bufio"
	"fmt"
	"os"
)

// 242. 有效的字母异位词
//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的 字母异位词。

//示例 1:
//输入: s = "anagram", t = "nagaram"
//输出: true

//示例 2:
//输入: s = "rat", t = "car"
//输出: false
//
//提示:
//1 <= s.length, t.length <= 5 * 104
//s 和 t 仅包含小写字母
//
//进阶: 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		s := scanner.Text()
		scanner.Scan()
		t := scanner.Text()
		fmt.Println(isAnagram(s, t))
	}
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hashmap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if s[i] == t[i] {
			continue
		}
		hashmap[s[i]]++
		hashmap[t[i]]--
	}

	for _, v := range hashmap {
		if v != 0 {
			return false
		}
	}
	return true
}
