package main

import (
	"container/list"
	"fmt"
)

func reverseWords(s string) string {
	//1.使用双指针删除冗余的空格
	slowIndex, fastIndex := 0, 0
	b := []byte(s)
	//删除头部冗余空格
	for len(b) > 0 && fastIndex < len(b) && b[fastIndex] == ' ' {
		fastIndex++
	}
	//删除单词间冗余空格
	for ; fastIndex < len(b); fastIndex++ {
		if fastIndex-1 > 0 && b[fastIndex-1] == b[fastIndex] && b[fastIndex] == ' ' {
			continue
		}
		b[slowIndex] = b[fastIndex]
		slowIndex++
	}
	//删除尾部冗余空格
	if slowIndex-1 > 0 && b[slowIndex-1] == ' ' {
		b = b[:slowIndex-1]
	} else {
		b = b[:slowIndex]
	}
	//2.反转整个字符串
	reverse(b, 0, len(b)-1)
	//3.反转单个单词  i单词开始位置，j单词结束位置
	i := 0
	for i < len(b) {
		j := i
		for ; j < len(b) && b[j] != ' '; j++ {
		}
		reverse(b, i, j-1)
		i = j
		i++
	}
	return string(b)
}

func reverse(b []byte, left, right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

// 利用stack先进后出的特点，先去掉开头结尾的空格，然后把单词逐个加入到stack里，最后遍历stack，拼接字符串
func reverseWordsByStack(s string) string {
	b := []byte(s)
	left := 0
	right := len(b) - 1
	// 去掉字符串开头的空白字符
	for left <= right && b[left] == ' ' {
		left++
	}

	// 去掉字符串末尾的空白字符
	for left <= right && b[right] == ' ' {
		right--
	}

	// fmt.Println(string(b[left : right+1]))

	stack := list.New()
	word := make([]byte, 0)

	for left <= right {
		c := b[left]
		if len(word) > 0 && c == ' ' {
			// 将单词 push 到队列的头部
			stack.PushBack(word)
			word = []byte("")
		} else if c != ' ' {
			word = append(word, c)
		}
		left++
	}
	stack.PushBack(word)

	ans := ""
	cur := stack.Back()
	for cur != nil {
		ans += string(cur.Value.([]byte))
		cur = cur.Prev()
		if cur != nil {
			ans += " "
		}
	}
	return ans
}

func main() {
	str := reverseWordsByStack("the sky is blue ")
	fmt.Println(str)
}
