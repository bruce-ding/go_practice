package main

import (
	"container/list"
	"fmt"
)

type CQueue struct {
	stack1, stack2 *list.List
}

func Constructor() CQueue {
	return CQueue{
		stack1: list.New(),
		stack2: list.New(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1.PushBack(value)
}

func (this *CQueue) DeleteHead() int {
	// 如果第二个栈为空
	if this.stack2.Len() == 0 {
		for this.stack1.Len() > 0 {
			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
		}
	}
	if this.stack2.Len() != 0 {
		e := this.stack2.Back()
		this.stack2.Remove(e)
		return e.Value.(int)
	}
	return -1
	//1. 如果B中有元素，则栈B中自栈顶到栈底，分别对应队列的队首往后依次配列。所以B中每次pop后，栈顶都是队列的队首。
	//2. 如果B中为空，则再从A中拿元素放入B中，然后再从B的栈顶删除元素即可。注意别忘了判断A是否为空。
	// if this.stack2.Len() != 0 {
	// 	e := this.stack2.Back()
	// 	this.stack2.Remove(e)
	// 	return e.Value.(int)
	// }
	// if this.stack1.Len() == 0 {
	// 	return -1
	// }
	// for this.stack1.Len() != 0 {
	// 	this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
	// }

	// e := this.stack2.Back()
	// this.stack2.Remove(e)
	// return e.Value.(int)
}

func main() {
	// arr := []string{"CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"}
	// [[],[],[5],[2],[],[]]
	// 输出：[null,-1,null,null,5,2]

	q := Constructor()
	fmt.Println(q.DeleteHead())
	q.AppendTail(5)
	fmt.Println(q.stack1)
	q.AppendTail(2)
	fmt.Println(q.stack1)
	fmt.Println(q.DeleteHead())
	q.AppendTail(9)
	fmt.Println(q.DeleteHead())
	fmt.Println(q.DeleteHead())

}
