package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//203. 移除链表元素
//给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
//
//
//示例 1：
//输入：head = [1,2,6,3,4,5,6], val = 6
//输出：[1,2,3,4,5]
//
//示例 2：
//输入：head = [], val = 1
//输出：[]
//
//示例 3：
//输入：head = [7,7,7,7], val = 7
//输出：[]
//
//提示：
//列表中的节点数目在范围 [0, 104] 内
//1 <= Node.val <= 50
//0 <= val <= 50
//

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var n int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		line := scanner.Text()
		n, _ = strconv.Atoi(line)
		var pre, head *ListNode
		for i := 0; i < n; i++ {
			node := &ListNode{}
			if pre == nil {
				pre = node
				head = node
			} else {
				pre.Next = node
			}
			scanner.Scan()
			line = scanner.Text()
			node.Val, _ = strconv.Atoi(line)
			pre = node
		}
		scanner.Scan()
		val, _ := strconv.Atoi(scanner.Text())
		for pre = head; pre != nil; pre = pre.Next {
			fmt.Print(pre.Val, " ")
		}

		fmt.Println(removeElements(head, val))
	}
}

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	pre := head
	var cur *ListNode
	if pre != nil {
		cur = pre.Next
	}
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return head
}
