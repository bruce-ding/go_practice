=begin
20. 有效的括号
力扣题目链接(https://leetcode-cn.com/problems/valid-parentheses/)

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:
输入: "()"
输出: true

示例 2:
输入: "()[]{}"
输出: true

示例 3:
输入: "(]"
输出: false

示例 4:
输入: "([)]"
输出: false

示例 5:
输入: "{[]}"
输出: true
=end

def is_valid(s)
	hash = {')' => '(', ']' => '[', '}' => '{'}
	stack = []
	return true if s == ""

	s.each_char do |c|
		if c == '(' || c == '[' || c == '{'
			stack.push c
		elsif stack.last == hash[c]
			stack.pop
		else
			return false
		end
	end

	stack.length == 0
end

require 'minitest/autorun'

describe "is valid parentheses" do
	it 'must' do
		str = "{[]}"
		assert_equal(true, is_valid(str))
		str = "([{}]()"
		assert_equal(false, is_valid(str))
		str = "([{}}}"
		assert_equal(false, is_valid(str))
		str = "([{}])))"
		assert_equal(false, is_valid(str))
	end
end


