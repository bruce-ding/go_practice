package main

import "fmt"

func reverseLeftWords(s string, n int) string {
	b := []byte(s)
	// 0. 若原始字符串       ab|cdefg
	// 1. 反转前n个字符      ba|cdefg
	// 2. 反转第n到end字符   ba|gfedc
	// 3. 反转整个字符       cdefg|ab
	reverse(b, 0, n-1)
	reverse(b, n, len(b)-1)
	reverse(b, 0, len(b)-1)
	return string(b)
}

// 切片是引用传递
func reverse(b []byte, left, right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

func main() {
	str := reverseLeftWords("abcdefg", 2)
	fmt.Println(str)
}
