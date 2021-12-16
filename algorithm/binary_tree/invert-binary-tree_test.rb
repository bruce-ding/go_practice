=begin
226. 翻转二叉树
力扣题目链接(https://leetcode-cn.com/problems/invert-binary-tree/)
翻转一棵二叉树。

示例：

输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9

输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1

=end

require_relative './binary_tree_node.rb'

# 前序遍历
def invert_tree(root)
	return root if root == nil

	# 先交换左右孩子
	root.left, root.right = root.right, root.left # 中
	# 再分别翻转左右子树
	invert_tree(root.left)  # 左
	invert_tree(root.right) # 右

	root
end

# 层序遍历
def invert_tree1(root)
	return root if root == nil

	queue = Queue.new
	queue.enq(root)

	while (len = queue.length) > 0
		len.times do
			node = queue.deq
			node.left, node.right = node.right, node.left #交换
			if node.left != nil
				queue.enq(e.left)
			end
			if e.right != nil
				queue.enq(node.right)
			end
		end
	end

	root
end

require 'minitest/autorun'

describe "invert_tree" do
	it "must" do
		root = TreeNode.new(1)
		root.left = TreeNode.new(2)
		root.left.left = TreeNode.new(4)
		root.left.right = TreeNode.new(5)
		root.right = TreeNode.new(3)
		root.right.left = TreeNode.new(6)
		root.right.right = TreeNode.new(7)
	
		fmt.Println("before invert")
		print_tree(root)
	
		res = invert_tree(root)
		fmt.Println("after invert")
		slice = print_tree(res)
		expected = [][]int{
			{1}, {3, 2}, {7, 6, 5, 4},
		}
		assert.Equal(t, expected, slice)
	
		res = invert_tree1(root)
		fmt.Println("invert back")
		slice = print_tree(res)
		expected = [][]int{
			{1}, {2, 3}, {4, 5, 6, 7},
		}
		assert_equal(expected, slice)
	end
end
