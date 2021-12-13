package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
题目：剑指Offer58-II.左旋转字符串
力扣题目链接(https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/)

字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。请定义一个函数实现字符串左旋转操作的功能。比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。

示例 1：
输入: s = "abcdefg", k = 2
输出: "cdefgab"

示例 2：
输入: s = "lrloseumgh", k = 6
输出: "umghlrlose"

限制：
1 <= k < s.length <= 10000
**/

func reverseLeftWords(s string, n int) string {
	b := []byte(s)
	// 0. 若原始字符串       ab|cdefg
	// 1. 反转前n个字符      ba|cdefg
	// 2. 反转第n到end字符   ba|gfedc
	// 3. 反转整个字符       cdefg|ab
	reverse1(b, 0, n-1)
	reverse1(b, n, len(b)-1)
	reverse1(b, 0, len(b)-1)
	return string(b)
}

// 切片是引用传递
func reverse1(bytes []byte, left, right int) {
	for left < right {
		bytes[left], bytes[right] = bytes[right], bytes[left]
		left++
		right--
	}
}

func TestReverseLeftWords(t *testing.T) {
	str := reverseLeftWords("abcdefg", 2)
	assert.Equal(t, "cdefgab", str)
}
