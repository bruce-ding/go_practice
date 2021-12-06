package hash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//242.有效的字母异位词
//力扣题目链接(https://leetcode-cn.com/problems/valid-anagram/)
//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
//注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。
//
//示例 1:
//
//输入: s = "anagram", t = "nagaram"
//输出: true
//示例 2:
//
//输入: s = "rat", t = "car"
//输出: false
//
//提示:
//
//1 <= s.length, t.length <= 5 * 104
//s 和 t 仅包含小写字母
//
//进阶: 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	for j := 0; j < len(t); j++ {
		_, ok := m[t[j]]
		if ok {
			m[t[j]]--
			if m[t[j]] < 0 {
				return false
			}
		} else {
			return false
		}

	}
	return true
}

func TestIsAnagram(t *testing.T) {
	res := isAnagram("anagram", "nagaram")
	assert.True(t, res, "res should be true")

	res1 := isAnagram("anagram", "abc")
	assert.False(t, res1, "res1 should be false")
}
