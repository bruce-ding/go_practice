# Definition for a binary tree node.
class TreeNode
    attr_accessor :val, :left, :right
    def initialize(val = 0, left = nil, right = nil)
        @val = val
        @left = left
        @right = right
    end
end

def print_tree(root)
    res = []
	return res if root == nil #防止为空
	
	queue = Queue.new
	queue.enq root

	while (len = queue.length) > 0
		level_arr = []
		len.times do
			node = queue.deq #出队列
			level_arr << node.val #将值加入本层数组中

			if node.left != nil
				queue.enq node.left
			end
			if node.right != nil
				queue.enq node.right
			end
		end
		res << level_arr #放入结果集
	end

	res
end