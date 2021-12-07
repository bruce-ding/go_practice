package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
题目：剑指Offer 05.替换空格
力扣题目链接(https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/)

请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

示例 1
输入：s = "We are happy."
输出："We%20are%20happy."
**/

// 遍历添加
func replaceSpace1(s string) string {
	b := []byte(s)
	result := make([]byte, 0)
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' {
			result = append(result, []byte("%20")...)
		} else {
			result = append(result, b[i])
		}
	}
	return string(result)
}

// 原地修改
func replaceSpace(s string) string {
	b := []byte(s)
	length := len(b)
	spaceCount := 0
	// 计算空格数量
	for _, v := range b {
		if v == ' ' {
			spaceCount++
		}
	}

	// 扩展原有切片
	resizeCount := spaceCount * 2
	tmp := make([]byte, resizeCount)
	b = append(b, tmp...)

	i := length - 1
	j := len(b) - 1
	for i >= 0 {
		if b[i] != ' ' {
			b[j] = b[i]
			i--
			j--
		} else {
			b[j] = '0'
			b[j-1] = '2'
			b[j-2] = '%'
			i--
			j -= 3
		}
	}

	return string(b)
}

func TestReplaceSpace(t *testing.T) {
	// 输入：s = "We are happy."
	// 输出："We%20are%20happy."
	str := replaceSpace("We are happy.")
	assert.Equal(t, "We%20are%20happy.", str)

	str = replaceSpace1("We are happy.")
	assert.Equal(t, "We%20are%20happy.", str)
}
