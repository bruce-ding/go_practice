package main

import "fmt"

// 题目链接：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/

var digitsMap = [10]string{
	"",     // 0
	"",     // 1
	"abc",  // 2
	"def",  // 3
	"ghi",  // 4
	"jkl",  // 5
	"mno",  // 6
	"pqrs", // 7
	"tuv",  // 8
	"wxyz", // 9
}

var result []string // 存放符合条件结果的集合
var s string        // 用来存放符合条件的结果

func backtracking(digits string, index int) {
	if index == len(digits) {
		result = append(result, s)
		return
	}
	digit := digits[index] - '0' // 将index指向的数字转为int
	letters := digitsMap[digit]  // 取数字对应的字符集
	for i := 0; i < len(letters); i++ {
		s += string(letters[i])       // 处理
		backtracking(digits, index+1) // 递归，注意index+1，一下层要处理下一个数字了
		s = s[:len(s)-1]              //回溯
	}
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return result
	}
	backtracking(digits, 0)
	return result
}

func main() {
	// 示例: 输入："23" 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]
	res := letterCombinations("23")
	fmt.Println(res)
}
