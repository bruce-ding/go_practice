package main

import "fmt"

func firstUniqChar(s string) byte {
	var list [26]int
	length := len(s)
	for i := 0; i < length; i++ {
		fmt.Printf("%T\n", s[i])
		list[s[i]-'a']++
	}
	for i := 0; i < length; i++ {
		if list[s[i]-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}

func main() {
	res := firstUniqChar("abacc")
	fmt.Print(res)
}
