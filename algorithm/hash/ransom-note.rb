# // 383. 赎金信
# // 力扣题目链接(https://leetcode-cn.com/problems/ransom-note/)
# // 为了不在赎金信中暴露字迹，从杂志上搜索各个需要的字母，组成单词来表达意思。
# // 给你一个赎金信 (ransomNote) 字符串和一个杂志(magazine)字符串，判断 ransomNote 能不能由 magazines 里面的字符构成。
# // 如果可以构成，返回 true ；否则返回 false 。
# // magazine 中的每个字符只能在 ransomNote 中使用一次。

# // 示例 1：
# // 输入：ransomNote = "a", magazine = "b"
# // 输出：false

# // 示例 2：
# // 输入：ransomNote = "aa", magazine = "ab"
# // 输出：false

# // 示例 3：
# // 输入：ransomNote = "aa", magazine = "aab"
# // 输出：true

# // 提示：
# // 1 <= ransomNote.length, magazine.length <= 105
# // ransomNote 和 magazine 由小写英文字母组成

def can_construct(ransom_note, magazine)
	record = []
	# codepoint of 'a' is 97
	magazine.each_codepoint do |c|
		i = c - 97
		record[i] = record[i].to_i + 1
	end
	ransom_note.each_codepoint do |c|
		i = c - 97
		record[i] = record[i].to_i - 1
		if record[i] < 0
			return false
		end
	end
	return true
end

def can_construct1(ransom_note, magzine)
	record = {}
	magzine.each_char do |c|
		record[c] = record[c].to_i + 1
	end

	ransom_note.each_char do |c|
		record[c] = record[c].to_i - 1
		if record[c] < 0
			return false
		end
	end

	true
end

res = can_construct("aa", "aab")
p(res) # should be true
res = can_construct1("aa", "aab")
p(res) # should be true
