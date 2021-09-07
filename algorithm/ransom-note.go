package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	record := make([]int, 26)
	for _, v := range magazine {
		record[v-'a']++
	}
	for _, v := range ransomNote {
		record[v-'a']--
		if record[v-'a'] < 0 {
			return false
		}
	}
	return true
}

func canConstruct1(ransomNote, magzine string) bool {
	record := make([]int, 26)
	for _, v := range magzine {
		record[v-'a']++
	}

	for _, v := range ransomNote {
		record[v-'a']--
		if record[v-'a'] < 0 {
			return false
		}
	}

	return true
}

func main() {
	res := canConstruct1("aa", "aab")
	fmt.Println(res)
}
