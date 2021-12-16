=begin
344.反转字符串
力扣题目链接(https://leetcode-cn.com/problems/reverse-string/)

编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。

不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。

你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。

示例 1：
输入：["h","e","l","l","o"]
输出：["o","l","l","e","h"]

示例 2：
输入：["H","a","n","n","a","h"]
输出：["h","a","n","n","a","H"]
=end

def reverse_string(s)
	left = 0
	right = s.length - 1
	while left < right
		s[left], s[right] = s[right], s[left]
		left += 1
		right -= 1
	end
end

def reverse_string1(s)
	l = s.length
	(0..l/2).each do |i|
		s[i], s[l-i-1] = s[l-i-1], s[i]
	end
end

require "minitest/autorun"

describe "reverse_string" do
	it "must" do
		# 输入：s = "abcdefg", k = 2
		# 输出："bacdfeg"
		s = "hello world"
		reverse_string(s)
		assert_equal s, "dlrow olleh"

		reverse_string1(s)
		assert_equal s, "hello world"
	end
end