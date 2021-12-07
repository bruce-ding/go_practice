=begin
题目：剑指Offer 05.替换空格
力扣题目链接(https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/)

请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

示例 1
输入：s = "We are happy."
输出："We%20are%20happy."
=end

# 遍历添加
def replace_space1(s)
	res = ''
	s.each_char do |c|
		if c == ' '
			res += '%20'
		else
			res += c
		end
	end
	
	res
end

# 原地修改
def replace_space(s)
	chars = s.chars
	length = chars.length
	# 计算空格数量
	space_count = chars.count{|c| c == ' '}
	
	# 扩展原有数组
	resize_count = space_count * 2
	chars += Array.new(resize_count, ' ')

	i = length - 1
	j = chars.length - 1
	while i >= 0
		if chars[i] != ' '
			chars[j] = chars[i]
			i -= 1
			j -= 1
		else
			chars[j] = '0'
			chars[j-1] = '2'
			chars[j-2] = '%'
			i -= 1
			j -= 3
		end
	end

	chars.join('')
end

# 输入：s = "We are happy."
# 输出："We%20are%20happy."
str = replace_space("We are happy.")
# p str # should be "We%20are%20happy."

str = replace_space1("We are happy.")
# p str # should be "We%20are%20happy."

require "minitest/autorun"

describe "replace_space" do
#   before do
#     @meme = Meme.new
#   end

  describe "replace space of \We are happy.\" do
    it "must return " do
		str = replace_space("We are happy.")
		_(str).must_equal "We%20are%20happy."
    end
  end

#   describe "when asked about blending possibilities" do
#     it "won't say no" do
#       _(@meme.will_it_blend?).wont_match /^no/i
#     end
#   end
end