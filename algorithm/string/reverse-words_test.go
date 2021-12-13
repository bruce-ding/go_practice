package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
151.翻转字符串里的单词
力扣题目链接(https://leetcode-cn.com/problems/reverse-words-in-a-string/)

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
**/

func reverseWords(s string) string {
	//1.使用双指针删除冗余的空格
	slow, fast := 0, 0
	bytes := []byte(s)
	length := len(bytes)
	//删除头部冗余空格
	for length > 0 && fast < length && bytes[fast] == ' ' {
		fast++
	}
	//删除单词间冗余空格
	for ; fast < length; fast++ {
		if fast-1 > 0 && bytes[fast-1] == bytes[fast] && bytes[fast] == ' ' {
			continue
		}
		bytes[slow] = bytes[fast]
		slow++
	}
	//删除尾部冗余空格
	if slow-1 > 0 && bytes[slow-1] == ' ' {
		length = slow - 1
	} else {
		length = slow
	}
	//2.反转整个字符串
	reverseBytes(bytes, 0, length-1)

	//3.反转单个单词  i单词开始位置，j单词结束位置
	i := 0
	for i < length {
		j := i
		for j < length && bytes[j] != ' ' {
			j++
		}
		reverseBytes(bytes, i, j-1)
		i = j
		i++
	}
	return string(bytes[0:length])
}

func reverseBytes(bytes []byte, left, right int) {
	for left < right {
		bytes[left], bytes[right] = bytes[right], bytes[left]
		left++
		right--
	}
}

func TestReplaceWords(t *testing.T) {
	// 	示例 1：
	// 输入: "the sky is blue"
	// 输出: "blue is sky the"
	str := reverseWords("the sky is blue ")
	assert.Equal(t, "blue is sky the", str)

	// 示例 2：
	// 输入: "  hello world!  "
	// 输出: "world! hello"
	// 解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
	str = reverseWords("  hello  world!  ")
	assert.Equal(t, "world! hello", str)

	// 示例 3：
	// 输入: "a good   example"
	// 输出: "example good a"
	// 解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
	str = reverseWords("a good   example")
	assert.Equal(t, "example good a", str)
}
