package main

// 206. 反转链表
//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
//
//
//示例 1：
//输入：head = [1,2,3,4,5]
//输出：[5,4,3,2,1]

//示例 2：
//输入：head = [1,2]
//输出：[2,1]

//示例 3：
//输入：head = []
//输出：[]
//
//提示：
//链表中节点的数目范围是 [0, 5000]
//-5000 <= Node.val <= 5000

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	dummyHead := &ListNode{}
	var head *ListNode = nil
	cur := dummyHead
	for i := 1; i < 6; i++ {
		cur.Next = &ListNode{Val: i}
		if head == nil {
			head = cur.Next
		}
		cur = cur.Next
	}
	reverseList(head)
}

func reverseList(head *ListNode) *ListNode {
	var ptr1 *ListNode = nil
	ptr2 := head
	if ptr2 == nil || ptr2.Next == nil {
		return ptr2
	}
	ptr3 := ptr2.Next
	for ptr2 != nil {
		ptr2.Next = ptr1
		ptr1 = ptr2
		ptr2 = ptr3
		if ptr3 != nil {
			ptr3 = ptr3.Next
		}
	}
	return ptr1
}
