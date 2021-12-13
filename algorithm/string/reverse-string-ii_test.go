package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
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
**/

func reverseStr(s string, k int) string {
	ss := []byte(s)
	length := len(s)
	for i := 0; i < length; i += 2 * k {
		if i+k <= length {
			reverse(ss[i : i+k])
		} else {
			reverse(ss[i:length])
		}
	}
	return string(ss)
}

func reverse(s []byte) {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func TestReverseStringII(t *testing.T) {
	// 输入：s = "abcdefg", k = 2
	// 输出："bacdfeg"
	str := reverseStr("abcdefg", 2)
	assert.Equal(t, "bacdfeg", str)

	str = reverseStr("abcdefghijklm", 3)
	assert.Equal(t, "cbadefihgjklm", str)
}
