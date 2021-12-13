=begin
给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。

如果剩余字符少于 k 个，则将剩余字符全部反转。
如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。

示例 1：
输入：s = "abcdefg", k = 2
输出："bacdfeg"

示例 2：
输入：s = "abcd", k = 2
输出："bacd"

提示：
1 <= s.length <= 104
s 仅由小写英文组成
1 <= k <= 104
=end

# @param {String} s
# @param {Integer} k
# @return {String}
def reverse_str(s, k)
	length = s.length
	res = ''
	i = 0
	while i < length
		if i % (2 * k) == 0
			if i+k <= length
				res += reverse(s[i..i+k-1])
			else
				res += reverse(s[i..length-1])
			end
			i += k
		else
			res += s[i]
			i += 1
		end
		
	end
	return res
end

def reverse(s)
	left = 0
	right = s.length - 1
	while left < right
		s[left], s[right] = s[right], s[left]
		left += 1
		right -= 1
	end
	s
end

require "minitest/autorun"

describe "reverse_str" do
	it "must" do
		# 输入：s = "abcdefg", k = 2
		# 输出："bacdfeg"
		str = reverse_str("abcdefg", 2)
		assert_equal str, "bacdfeg"

		str = reverse_str("abcdefghijklm", 3)
		assert_equal str, "cbadefihgjklm"
	end
end