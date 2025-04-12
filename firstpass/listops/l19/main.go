package main

//19. 删除链表的倒数第 N 个结点
//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
//示例 1：
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]

//示例 2：
//输入：head = [1], n = 1
//输出：[]

//示例 3：
//输入：head = [1,2], n = 1
//输出：[1]
//
//提示：
//
//链表中结点的数目为 sz
//1 <= sz <= 30
//0 <= Node.val <= 100
//1 <= n <= sz

type ListNode struct {
	Next *ListNode
	Val  int
}

func main() {
	dummyHead := &ListNode{}
	for i := 1; i <= 5; i++ {
		dummyHead.Next = &ListNode{Val: i}
	}
	removeNthFromEnd(dummyHead.Next, 2)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{head, 0}
	slow, fast := dummyHead, dummyHead
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummyHead.Next
}
