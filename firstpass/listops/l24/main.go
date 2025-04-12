package main

// 24. 两两交换链表中的节点
// 中等
// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
//
// 示例 1：
// 输入：head = [1,2,3,4]
// 输出：[2,1,4,3]
//
// 示例 2：
// 输入：head = []
// 输出：[]
//
// 示例 3：
// 输入：head = [1]
// 输出：[1]
//
// 提示：
// 链表中节点的数目在范围 [0, 100] 内
// 0 <= Node.val <= 100
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	dummyHead := &ListNode{}
	var head *ListNode = nil
	for i, cur := 1, dummyHead; i < 5; i++ {
		cur.Next = &ListNode{Val: i}
		if head == nil {
			head = cur.Next
		}
		cur = cur.Next
	}
	swapPairs(head)
}

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	ptr1, ptr2 := dummyHead, dummyHead.Next
	if ptr2 == nil || ptr2.Next == nil {
		return dummyHead.Next
	}
	ptr3 := ptr2.Next
	for ptr2 != nil && ptr3 != nil {
		ptr2.Next = ptr3.Next
		ptr3.Next = ptr2
		ptr1.Next = ptr3

		ptr1 = ptr2
		ptr2 = ptr1.Next
		if ptr2 != nil {
			ptr3 = ptr2.Next
		}
	}
	return dummyHead.Next
}
