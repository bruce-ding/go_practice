package linked_list

import (
	"fmt"
	"testing"
)

// https://leetcode-cn.com/problems/design-linked-list/
// 707. 设计链表
// 设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，next 是指向下一个节点的指针/引用。如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。

// 在链表类中实现这些功能：

// get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
// addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
// addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
// addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
// deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。

// 示例：

// MyLinkedList linkedList = new MyLinkedList();
// linkedList.addAtHead(1);
// linkedList.addAtTail(3);
// linkedList.addAtIndex(1,2);   //链表变为1-> 2-> 3
// linkedList.get(1);            //返回2
// linkedList.deleteAtIndex(1);  //现在链表是1-> 3
// linkedList.get(1);            //返回3

type MyLinkedList struct {
	dummyHead *ListNode
	Size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		dummyHead: &ListNode{Val: -1}, // 这里定义的头结点 是一个虚拟头结点，而不是真正的链表头结点
		Size: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index > this.Size -1 {
		return -1
	}

	curNode := this.dummyHead.Next
	for index > 0 {
		curNode = curNode.Next
		index--
	}
	return curNode.Val
}

func (this *MyLinkedList) AddAtHead(val int)  {
	nextHead := this.dummyHead.Next
	this.dummyHead.Next = &ListNode{Val: val}
	this.dummyHead.Next.Next = nextHead
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int)  {
	curHead := this.dummyHead
	for curHead.Next != nil {
		curHead = curHead.Next
	}
	curHead.Next = &ListNode{Val: val}
	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index > this.Size {
		return
	}

	curNode := this.dummyHead
	for index > 0 {
		curNode = curNode.Next
		index--
	}
	newNode := &ListNode{Val: val}
	newNode.Next = curNode.Next
	curNode.Next = newNode
	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index < 0 || index > this.Size -1 {
		return
	}
	curNode := this.dummyHead

	for index > 0 {
		curNode = curNode.Next
		index--
	}

	curNode.Next = curNode.Next.Next
	this.Size--
}

func (this *MyLinkedList) PrintList() {
	curNode := this.dummyHead.Next
	for curNode != nil {
		fmt.Printf("%d -> ", curNode.Val)
		curNode = curNode.Next
	}
	fmt.Println()
}

func TestMyLinkedListTest(t *testing.T) {
	//  * Your MyLinkedList object will be instantiated and called as such:
	linkedList := Constructor()

	linkedList.AddAtHead(1)  //链表变为1->
	linkedList.PrintList()

	linkedList.AddAtTail(3)  //链表变为1-> 3
	linkedList.PrintList()

	linkedList.AddAtIndex(1,2)   //链表变为1-> 2-> 3
	linkedList.PrintList()

	val := linkedList.Get(1)          //返回2
	fmt.Println(val)

	linkedList.DeleteAtIndex(1)  //现在链表是1-> 3
	linkedList.PrintList()

	val = linkedList.Get(1)      //返回3
	fmt.Println(val)
}
