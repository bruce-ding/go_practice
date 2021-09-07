package main

import "fmt"

// 题目链接：https://leetcode-cn.com/problems/palindrome-partitioning/

// 注意切片（go切片是披着值类型外衣的引用类型）
var result [][]string
var path []string // 放已经回文的子串

func backtracking(s string, startIndex int) {
	// 如果起始位置已经大于s的大小，说明已经找到了一组分割方案了
	if startIndex >= len(s) {
		temp := make([]string, len(path))
		copy(temp, path)
		result = append(result, temp) // 处理节点
		return
	}

	for i := startIndex; i < len(s); i++ {
		if isPalindrome(s, startIndex, i) { // 是回文子串
			// 获取[startIndex,i]在s中的子串
			str := s[startIndex : i+1]
			path = append(path, str)
		} else { // 不是回文，跳过
			continue
		}
		backtracking(s, i+1)      // 寻找i+1为起始位置的子串
		path = path[:len(path)-1] // 回溯过程，弹出本次已经填在的子串
	}
}

func isPalindrome(s string, start, end int) bool {
	left := start
	right := end
	for left < right {
		if s[left] != s[right] {
			return false
		}
		//移动左右指针
		left++
		right--
	}
	return true
}

func partition(s string) [][]string {
	backtracking(s, 0)
	return result
}

func main() {
	// 示例: 输入: "aab" 输出: [ ["aa","b"], ["a","a","b"] ]
	res := partition("aab")
	fmt.Println(res)

}
