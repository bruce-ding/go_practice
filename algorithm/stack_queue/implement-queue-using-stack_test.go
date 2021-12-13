package stack_queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
232.用栈实现队列
力扣题目链接(https://leetcode-cn.com/problems/implement-queue-using-stacks/)

使用栈实现队列的下列操作：

push(x) -- 将一个元素放入队列的尾部。
pop() -- 从队列首部移除元素。
peek() -- 返回队列首部的元素。
empty() -- 返回队列是否为空。

示例:

MyQueue queue = new MyQueue();
queue.push(1);
queue.push(2);
queue.peek();  // 返回 1
queue.pop();   // 返回 1
queue.empty(); // 返回 false
说明:

你只能使用标准的栈操作 -- 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）。
#思路
这是一道模拟题，不涉及到具体算法，考察的就是对栈和队列的掌握程度。

使用栈来模式队列的行为，如果仅仅用一个栈，是一定不行的，所以需要两个栈一个输入栈，一个输出栈，这里要注意输入栈和输出栈的关系。
**/

type MyQueue struct {
	stackIn  []int
	stackOut []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		stackIn:  make([]int, 0),
		stackOut: make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stackIn = append(this.stackIn, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	// 只有当stackOut为空的时候，再从stackIn里导入数据（导入stackIn全部数据）
	if len(this.stackOut) == 0 {
		// 从stIn导入数据直到stIn为空
		for len(this.stackIn) != 0 {
			top := this.stackIn[len(this.stackIn)-1]
			this.stackOut = append(this.stackOut, top)
			this.stackIn = this.stackIn[:len(this.stackIn)-1]
		}
	}
	result := this.stackOut[len(this.stackOut)-1]
	this.stackOut = this.stackOut[:len(this.stackOut)-1]
	return result
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	res := this.Pop()                          // 直接使用已有的pop函数
	this.stackOut = append(this.stackOut, res) // 因为pop函数弹出了元素res，所以再添加回去
	return res
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stackIn) == 0 && len(this.stackOut) == 0
}

func TestMyQueue(t *testing.T) {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	assert.Equal(t, 1, queue.Peek()) // 返回 1
	assert.Equal(t, 1, queue.Pop())  // 返回 1
	assert.False(t, queue.Empty())   // 返回 false
}
