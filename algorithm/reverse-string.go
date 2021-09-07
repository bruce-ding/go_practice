package main

import (
	"fmt"
	"strings"
)

func reverseString(s []byte) {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func reverseString1(s []byte) {
	l := len(s)
	for i := 0; i < l; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
}

func main() {
	bytes := []byte("hello world")
	reverseString1(bytes)
	fmt.Println(bytes)
	for _, v := range bytes {
		fmt.Println(string(v))
	}
	strings.Replace("abc", "a", "d", -1)
}
