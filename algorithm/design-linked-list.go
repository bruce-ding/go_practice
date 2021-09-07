package main

import "fmt"

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

type Node struct {
	val  int
	next *Node
	prev *Node
}

type MyLinkedList struct {
	head *Node
	tail *Node
	val  int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{head: nil, tail: nil, val: 0}

}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	head := this.head
	i := 0
	for head != nil {
		if i == index {
			return head.val
		}
		i = i + 1
		head = head.next
	}
	return -1
}

//打印链表
func (this *MyLinkedList) PrintList() {
	head := this.head
	for head != nil {
		fmt.Print(head.val, "->")
		head = head.next
	}
}

// 1->2

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	head := this.head
	newHead := &Node{next: head, prev: nil, val: val}
	this.head = newHead
	if head != nil {
		head.prev = newHead
	} else {
		this.tail = newHead
	}
	this.val = this.val + 1
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	tail := this.tail
	newTail := &Node{next: nil, prev: tail, val: val}
	if this.tail != nil {
		this.tail.next = newTail
	} else {
		this.head = newTail
	}
	this.tail = newTail
	this.val = this.val + 1
}

/**
 */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	i := 0
	newNode := &Node{prev: nil, next: nil, val: val}
	cur := this.head
	switch {
	//增加头结点
	case index <= 0:
		this.AddAtHead(val)
	//增加尾结点
	case index == this.val:
		this.AddAtTail(val)
	//增加中间节点
	default:
		for cur != nil {
			if i == index {
				tmp := cur.prev
				tmp.next = newNode
				newNode.next = cur
				newNode.prev = tmp
				cur.prev = newNode
				this.val = this.val + 1
			}
			i = i + 1
			cur = cur.next
		}
	}

}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	cur := this.head
	i := 0
	//删除头结点
	if index == 0 {
		if cur.next != nil {
			this.head = cur.next
		} else {
			this.head = nil
		}
		this.val = this.val - 1
		return
	}
	if index == this.val-1 {
		//删除尾结点
		fmt.Println(this.tail.val, this.tail.prev.val)
		tmp := this.tail.prev
		fmt.Println(tmp.val)
		tmp.next = nil
		this.tail = tmp
		this.val = this.val - 1
		return
	} else {
		//删除中间节点
		for cur != nil {
			if i == index {
				tmp := cur.prev
				tmp.next = cur.next
				cur.next.prev = tmp
				this.val = this.val - 1
				return
			} else {
				i = i + 1
				cur = cur.next
			}
		}
	}
}

func main() {

	//  * Your MyLinkedList object will be instantiated and called as such:
	index := 3
	val := 4
	obj := Constructor()
	obj.head = &Node{val: 1}
	obj.head.next = &Node{val: 2}
	obj.head.next.next = &Node{val: 3}
	obj.head.next.next.next = &Node{val: 4}
	obj.head.next.next.next.next = &Node{val: 5}
	element := obj.Get(index)
	fmt.Println(element)
	obj.AddAtHead(val)
	obj.AddAtTail(val)
	obj.AddAtIndex(index, val)
	obj.DeleteAtIndex(index)

}
