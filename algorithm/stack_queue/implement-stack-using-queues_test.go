/**
225. 用队列实现栈
力扣题目链接(https://leetcode-cn.com/problems/implement-stack-using-queues/)

使用队列实现栈的下列操作：

push(x) -- 元素 x 入栈
pop() -- 移除栈顶元素
top() -- 获取栈顶元素
empty() -- 返回栈是否为空
注意:

你只能使用队列的基本操作-- 也就是 push to back, peek/pop from front, size, 和 is empty 这些操作是合法的。
你所使用的语言也许不支持队列。 你可以使用 list 或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。
你可以假设所有操作都是有效的（例如, 对一个空的栈不会调用 pop 或者 top 操作）
**/
package stack_queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MyStack struct {
	queue []int //创建一个队列
}

/** Initialize your data structure here. */
func Constructor1() MyStack {
	return MyStack{ //初始化
		queue: make([]int, 0),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	//添加元素
	this.queue = append(this.queue, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	n := len(this.queue) - 1 //判断长度
	for n != 0 {             //除了最后一个，其余的都重新添加到队列里
		val := this.queue[0]
		this.queue = this.queue[1:]
		this.queue = append(this.queue, val)
		n--
	}
	//弹出元素
	val := this.queue[0]
	this.queue = this.queue[1:]
	return val

}

/** Get the top element. */
func (this *MyStack) Top() int {
	//利用Pop函数，弹出来的元素重新添加
	val := this.Pop()
	this.queue = append(this.queue, val)
	return val
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

/**
示例：

输入：
["MyStack", "push", "push", "top", "pop", "empty"]
[[], [1], [2], [], [], []]
输出：
[null, null, null, 2, 2, false]

解释：
MyStack myStack = new MyStack();
myStack.push(1);
myStack.push(2);
myStack.top(); // 返回 2
myStack.pop(); // 返回 2
myStack.empty(); // 返回 False
**/

func TestMyStack(t *testing.T) {
	stack := Constructor1()
	stack.Push(1)
	stack.Push(2)
	assert.Equal(t, 2, stack.Top()) // 返回 2
	assert.Equal(t, 2, stack.Pop()) // 返回 2
	assert.False(t, stack.Empty())  // 返回 false
}
