package main

import (
	"container/list"
	"fmt"
)

type MinStack struct {
	dataStack *list.List
	minStack  *list.List
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		dataStack: list.New(),
		minStack:  list.New(),
	}
}

func (this *MinStack) Push(x int) {
	this.dataStack.PushBack(x)
	if this.minStack.Len() == 0 {
		this.minStack.PushBack(x)
	} else {
		backElement := this.minStack.Back().Value.(int)
		if backElement > x {
			this.minStack.PushBack(x)
		} else {
			this.minStack.PushBack(backElement)
		}
	}
}

func (this *MinStack) Pop() {
	this.dataStack.Remove(this.dataStack.Back())
	this.minStack.Remove(this.minStack.Back())
}

func (this *MinStack) Top() int {
	return this.dataStack.Back().Value.(int)
}

func (this *MinStack) Min() int {
	return this.minStack.Back().Value.(int)
}

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.Min()) // --> 返回 -3.
	minStack.Pop()
	minStack.Pop()
	fmt.Println(minStack.Min()) // --> 返回 -2.

}
