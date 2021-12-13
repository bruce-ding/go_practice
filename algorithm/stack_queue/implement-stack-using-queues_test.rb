=begin
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
=end
class MyStack
	attr_accessor :queue
    def initialize()
		@queue = Queue.new
    end

=begin
    :type x: Integer
    :rtype: Void
=end
    def push(x)
		@queue.enq x
    end

=begin
    :rtype: Integer
=end
    def pop()
		n = @queue.length - 1
		while n > 0
			val = @queue.deq
			@queue.enq val
			n -= 1
		end
		@queue.deq
    end

=begin
    :rtype: Integer
=end
    def top()
		val = self.pop()
		@queue.enq val
		val
    end

=begin
    :rtype: Boolean
=end
    def empty()
		queue.length == 0
    end

end

=begin
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
=end

require 'minitest/autorun'

describe "MyStack" do
	it "must" do
		stack = MyStack.new

		stack.push(1)
		stack.push(2)
		assert_equal(2, stack.top()) # 返回 2
		assert_equal(2, stack.pop()) # 返回 2
		assert_equal(false, stack.empty())  # 返回 false
	end
end



