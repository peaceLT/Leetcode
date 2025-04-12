package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//707. 设计链表
//你可以选择使用单链表或者双链表，设计并实现自己的链表。
//
//单链表中的节点应该具备两个属性：val 和 next 。val 是当前节点的值，next 是指向下一个节点的指针/引用。
//
//如果是双向链表，则还需要属性 prev 以指示链表中的上一个节点。假设链表中的所有节点下标从 0 开始。
//
//实现 MyLinkedList 类：
//
//MyLinkedList() 初始化 MyLinkedList 对象。
//int get(int index) 获取链表中下标为 index 的节点的值。如果下标无效，则返回 -1 。
//void addAtHead(int val) 将一个值为 val 的节点插入到链表中第一个元素之前。在插入完成后，新节点会成为链表的第一个节点。
//void addAtTail(int val) 将一个值为 val 的节点追加到链表中作为链表的最后一个元素。
//void addAtIndex(int index, int val) 将一个值为 val 的节点插入到链表中下标为 index 的节点之前。如果 index 等于链表的长度，那么该节点会被追加到链表的末尾。如果 index 比长度更大，该节点将 不会插入 到链表中。
//void deleteAtIndex(int index) 如果下标有效，则删除链表中下标为 index 的节点。
//
//
//示例：
//
//输入
//["MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get", "deleteAtIndex", "get"]
//[[], [1], [3], [1, 2], [1], [1], [1]]
//输出
//[null, null, null, null, 2, null, 3]
//
//解释
//MyLinkedList myLinkedList = new MyLinkedList();
//myLinkedList.addAtHead(1);
//myLinkedList.addAtTail(3);
//myLinkedList.addAtIndex(1, 2);    // 链表变为 1->2->3
//myLinkedList.get(1);              // 返回 2
//myLinkedList.deleteAtIndex(1);    // 现在，链表变为 1->3
//myLinkedList.get(1);              // 返回 3
//
//
//提示：
//
//0 <= index, val <= 1000
//请不要使用内置的 LinkedList 库。
//调用 get、addAtHead、addAtTail、addAtIndex 和 deleteAtIndex 的次数不超过 2000 。

// 处理输入的函数
func processInput(operations []string, params [][]int) []interface{} {
	var result []interface{}
	var list MyLinkedList
	for i, op := range operations {
		switch op {
		case "MyLinkedList":
			list = Constructor()
			result = append(result, nil)
		case "addAtHead":
			list.AddAtHead(params[i][0])
			result = append(result, nil)
		case "addAtTail":
			list.AddAtTail(params[i][0])
			result = append(result, nil)
		case "addAtIndex":
			list.AddAtIndex(params[i][0], params[i][1])
			result = append(result, nil)
		case "get":
			res := list.Get(params[i][0])
			result = append(result, res)
		case "deleteAtIndex":
			list.DeleteAtIndex(params[i][0])
			result = append(result, nil)
		}
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// 读取操作字符串
	fmt.Println("请输入操作字符串（例如：[\"MyLinkedList\", \"addAtHead\", \"addAtTail\"]）：")
	scanner.Scan()
	operationStr := scanner.Text()
	//operationStr = strings.TrimPrefix(operationStr, "[")
	//operationStr = strings.TrimSuffix(operationStr, "]")
	operationStrs := strings.Split(operationStr, ",")
	operations := make([]string, 0, len(operationStrs))
	for _, op := range operationStrs {
		op = strings.TrimSpace(op)
		//op = strings.Trim(op, "\"")
		operations = append(operations, op)
	}

	// 读取参数
	fmt.Println("请输入参数（例如：[], [1], [3]）：")
	scanner.Scan()
	paramStr := scanner.Text()
	//paramStr = strings.TrimPrefix(paramStr, "[")
	//paramStr = strings.TrimSuffix(paramStr, "]")
	paramStrs := strings.Split(paramStr, "],")
	params := make([][]int, 0, len(paramStrs))
	for _, pStr := range paramStrs {
		pStr = strings.TrimSpace(pStr)
		pStr = strings.Trim(pStr, "[]")
		if pStr == "" {
			params = append(params, []int{})
			continue
		}
		pNums := strings.Split(pStr, ",")
		nums := make([]int, 0, len(pNums))
		for _, numStr := range pNums {
			numStr = strings.TrimSpace(numStr)
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}
		params = append(params, nums)
	}

	output := processInput(operations, params)
	fmt.Println(output)
}

// 单链表实现
type SingleNode struct {
	Val  int         // 节点的值
	Next *SingleNode // 下一个节点的指针
}

type MyLinkedList struct {
	dummyHead *SingleNode // 虚拟头节点
	Size      int         // 链表大小
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	newNode := &SingleNode{ // 创建新节点
		-999,
		nil,
	}
	return MyLinkedList{ // 返回链表
		dummyHead: newNode,
		Size:      0,
	}

}

/*
  - Get the value of the index-th node in the linked list. If the index is
    invalid, return -1.
*/
func (this *MyLinkedList) Get(index int) int {
	/*if this != nil || index < 0 || index > this.Size {
		return -1
	}*/
	if this == nil || index < 0 || index >= this.Size { // 如果索引无效则返回-1
		return -1
	}
	// 让cur等于真正头节点
	cur := this.dummyHead.Next   // 设置当前节点为真实头节点
	for i := 0; i < index; i++ { // 遍历到索引所在的节点
		cur = cur.Next
	}
	return cur.Val // 返回节点值
}

/*
  - Add a node of value val before the first element of the linked list. After
    the insertion, the new node will be the first node of the linked list.
*/
func (this *MyLinkedList) AddAtHead(val int) {
	// 以下两行代码可用一行代替
	// newNode := new(SingleNode)
	// newNode.Val = val
	newNode := &SingleNode{Val: val}   // 创建新节点
	newNode.Next = this.dummyHead.Next // 新节点指向当前头节点
	this.dummyHead.Next = newNode      // 新节点变为头节点
	this.Size++                        // 链表大小增加1
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &SingleNode{Val: val} // 创建新节点
	cur := this.dummyHead            // 设置当前节点为虚拟头节点
	for cur.Next != nil {            // 遍历到最后一个节点
		cur = cur.Next
	}
	cur.Next = newNode // 在尾部添加新节点
	this.Size++        // 链表大小增加1
}

/*
  - Add a node of value val before the index-th node in the linked list. If
    index equals to the length of linked list, the node will be appended to the
    end of linked list. If index is greater than the length, the node will not be
    inserted.
*/
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 { // 如果索引小于0，设置为0
		index = 0
	} else if index > this.Size { // 如果索引大于链表长度，直接返回
		return
	}

	newNode := &SingleNode{Val: val} // 创建新节点
	cur := this.dummyHead            // 设置当前节点为虚拟头节点
	for i := 0; i < index; i++ {     // 遍历到指定索引的前一个节点
		cur = cur.Next
	}
	newNode.Next = cur.Next // 新节点指向原索引节点
	cur.Next = newNode      // 原索引的前一个节点指向新节点
	this.Size++             // 链表大小增加1
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size { // 如果索引无效则直接返回
		return
	}
	cur := this.dummyHead        // 设置当前节点为虚拟头节点
	for i := 0; i < index; i++ { // 遍历到要删除节点的前一个节点
		cur = cur.Next
	}
	if cur.Next != nil {
		cur.Next = cur.Next.Next // 当前节点直接指向下下个节点，即删除了下一个节点
	}
	this.Size-- // 注意删除节点后应将链表大小减一
}

// 打印链表
func (list *MyLinkedList) printLinkedList() {
	cur := list.dummyHead // 设置当前节点为虚拟头节点
	for cur.Next != nil { // 遍历链表
		fmt.Println(cur.Next.Val) // 打印节点值
		cur = cur.Next            // 切换到下一个节点
	}
}
