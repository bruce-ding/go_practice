package main

import (
	"fmt"
	"math/rand"
)

func shuffle(src []string) []string {
	dest := make([]string, len(src))
	perm := rand.Perm(len(src))
	for i, v := range perm {
		dest[v] = src[i]
	}
	return dest
}

func main() {
	strs := []string{"a", "b", "c", "d", "e"}
	res := shuffle(strs)
	fmt.Println(res)
}
