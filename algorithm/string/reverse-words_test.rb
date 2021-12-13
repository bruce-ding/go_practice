=begin
151.翻转字符串里的单词
力扣题目链接(https:#leetcode-cn.com/problems/reverse-words-in-a-string/)

给定一个字符串，逐个翻转字符串中的每个单词。

示例 1：
输入: "the sky is blue"
输出: "blue is sky the"

示例 2：
输入: "  hello world!  "
输出: "world! hello"
解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。

示例 3：
输入: "a good   example"
输出: "example good a"
解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
=end

def reverse_words(s)
	# 1.使用双指针删除冗余的空格
	slow, fast = 0, 0
	chars = s.chars
	len = chars.length
	# 删除头部冗余空格
	while len > 0 && fast < len && chars[fast] == ' '
		fast += 1
	end

	while fast < len
		if fast-1 > 0 && chars[fast-1] == chars[fast] && chars[fast] == ' '
			fast += 1
			next
		end
		chars[slow] = chars[fast]
		slow += 1
		fast += 1
	end

	# 删除尾部冗余空格
	if slow-1 > 0 && chars[slow-1] == ' '
		len = slow - 1
	else
		len = slow
	end
	
	# 2.反转整个字符串
	reverse_chars(chars, 0, len-1)
	# 3.反转单个单词  i单词开始位置，j单词结束位置
	i = 0
	while i < len
		j = i
		while j < len && chars[j] != ' '
			j += 1
		end
		reverse_chars(chars, i, j-1)
		i = j
		i += 1
	end
	return chars[0..len-1].join('')
end

def reverse_chars(chars, left, right)
	while left < right
		chars[left], chars[right] = chars[right], chars[left]
		left += 1
		right -= 1
	end
end

require "minitest/autorun"

describe 'reverse_words' do
	it "must" do
		# 示例 1：
		# 输入: "the sky is blue"
		# 输出: "blue is sky the"
		str = reverse_words("the sky is blue ")
		assert_equal("blue is sky the", str)

		# 示例 2：
		# 输入: "  hello world!  "
		# 输出: "world! hello"
		# 解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
		# 以下是变化过程
		# "  hello  world!  " slow = 0, fast = 0
		# "hello  world!  "   slow = 0, fast = 2
		# "hello world! "     slow = 11
		# "hello world!"
		# "world! hello"
		str = reverse_words("  hello  world!  ")
		assert_equal("world! hello", str)

		# 示例 3：
		# 输入: "a good   example"
		# 输出: "example good a"
		# 解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
		str = reverse_words("a good   example")
		assert_equal("example good a", str)
	end
end

