=begin
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
=end

class MyQueue
	attr_accessor :stack_in, :stack_out
    def initialize()
		@stack_in = []
		@stack_out = []
    end

=begin
    :type x: Integer
    :rtype: Void
=end
    def push(x)
		@stack_in.push x
    end

=begin
    :rtype: Integer
=end
    def pop()
		if @stack_out.length == 0
			while @stack_in.length != 0
                @stack_out.push(@stack_in.pop())
            end
		end
		@stack_out.pop
    end

=begin
    :rtype: Integer
=end
    def peek()
		x = self.pop()
		@stack_out.push x
		x
    end

=begin
    :rtype: Boolean
=end
    def empty()
		@stack_in.length == 0 && @stack_out.length == 0
    end

end

require 'minitest/autorun'

describe "myQueue" do
	it "must" do
		# Your MyQueue object will be instantiated and called as such:
		queue = MyQueue.new()

		p queue.stack_in
		p queue.stack_out

		queue.push(1)
		p queue.stack_in
		p queue.stack_out
		queue.push(2)
		p queue.stack_in
		p queue.stack_out
		assert_equal(1, queue.peek()) # 返回 1
		p queue.stack_in
		p queue.stack_out
		assert_equal(1, queue.pop())  # 返回 1
		p queue.stack_in
		p queue.stack_out
		assert_equal(false, queue.empty()) # 返回 false
	end
end

