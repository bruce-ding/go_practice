package main

import "fmt"

// 1641. 统计字典序元音字符串的数目
// 给你一个整数 n，请返回长度为 n 、仅由元音 (a, e, i, o, u) 组成且按 字典序排列 的字符串数量。

// 字符串 s 按 字典序排列 需要满足：对于所有有效的 i，s[i] 在字母表中的位置总是与 s[i+1] 相同或在 s[i+1] 之前。

// 示例 1：

// 输入：n = 1
// 输出：5
// 解释：仅由元音组成的 5 个字典序字符串为 ["a","e","i","o","u"]
// 示例 2：

// 输入：n = 2
// 输出：15
// 解释：仅由元音组成的 15 个字典序字符串为
// ["aa","ae","ai","ao","au","ee","ei","eo","eu","ii","io","iu","oo","ou","uu"]
// 注意，"ea" 不是符合题意的字符串，因为 'e' 在字母表中的位置比 'a' 靠后
// 示例 3：

// 输入：n = 33
// 输出：66045

func countVowelStrings1(n int) int {
	dp := make([][5]int, n+1)

	//初始化n=1的情况
	for i := 0; i < 5; i++ {
		dp[1][i] = 1
	}

	for i := 2; i <= n; i++ {
		//长度i的以u结尾的字符串可以由任意一个长度i-1的字符串结尾加个u得到
		dp[i][4] = dp[i-1][0] + dp[i-1][1] + dp[i-1][2] + dp[i-1][3] + dp[i-1][4]
		dp[i][3] = dp[i-1][0] + dp[i-1][1] + dp[i-1][2] + dp[i-1][3]
		dp[i][2] = dp[i-1][0] + dp[i-1][1] + dp[i-1][2]
		dp[i][1] = dp[i-1][0] + dp[i-1][1]
		//长度i的以a结尾的字符串只能由长度i-1的以a结尾的字符串结尾加个a得到
		dp[i][0] = dp[i-1][0]
	}

	//最终答案求个和就行啦
	return dp[n][0] + dp[n][1] + dp[n][2] + dp[n][3] + dp[n][4]
}

func countVowelStrings(n int) int {
	dp := map[rune]int{'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1}
	for i := 2; i <= n; i++ {
		dp = map[rune]int{
			'a': dp['a'],
			'e': dp['a'] + dp['e'],
			'i': dp['a'] + dp['e'] + dp['i'],
			'o': dp['a'] + dp['e'] + dp['i'] + dp['o'],
			'u': dp['a'] + dp['e'] + dp['i'] + dp['o'] + dp['u'],
		}
	}
	return dp['a'] + dp['e'] + dp['i'] + dp['o'] + dp['u']
}

func main() {
	res := countVowelStrings1(5)
	fmt.Println(res)
}
