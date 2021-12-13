package stack_queue

/**
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
**/

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isValid(s string) bool {
	m := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := make([]byte, 0)
	if s == "" {
		return true
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 && stack[len(stack)-1] == m[s[i]] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

func TestIsValidParentheses(t *testing.T) {
	str := "{[]}"
	assert.Equal(t, true, isValid(str))
	str = "([{}]()"
	assert.Equal(t, false, isValid(str))
	str = "([{}}}"
	assert.Equal(t, false, isValid(str))
	str = "([{}])))"
	assert.Equal(t, false, isValid(str))
}
