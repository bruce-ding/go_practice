# 二叉树的遍历

require_relative './binary_tree_node.rb'

def pre_order_traversal(root)
	return if root == nil

	p(root.val)
	pre_order_traversal(root.left)
	pre_order_traversal(root.right)
end

def mid_order_traversal(root)
	return if root == nil

	mid_order_traversal(root.left)
	p(root.val)
	mid_order_traversal(root.right)
end

def post_order_traversal(root)
	return if root == nil

	post_order_traversal(root.left)
	post_order_traversal(root.right)
	p(root.val)
end

def layer_order_traversal(root)
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

require 'minitest/autorun'
describe "traversal" do
	it "must" do
		root = TreeNode.new(5)
		root.left = TreeNode.new(4)
		root.left.left = TreeNode.new(1)
		root.left.right = TreeNode.new(2)
		root.left.right.left = TreeNode.new(9)
		root.left.right.right = TreeNode.new(10)
		root.right = TreeNode.new(6)
		root.right.left = TreeNode.new(7)
		root.right.right = TreeNode.new(8)

		pre_order_traversal(root)
		mid_order_traversal(root)
		post_order_traversal(root)

		res = layer_order_traversal(root)
		expected = [
			[5],
			[4, 6],
			[1, 2, 7, 8],
			[9, 10]
			]
		assert_equal(expected, res)
	end
end
