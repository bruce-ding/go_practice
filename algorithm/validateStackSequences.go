package main

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
	//建一个辅助栈
	stack := make([]int, 0)
	//将push顺序压入辅助栈中，如果栈顶元素==pop序列中下一个出现的值，则弹出
	i := 0
	for _, value := range pushed {
		stack = append(stack, value)
		for len(stack) != 0 && stack[len(stack)-1] == popped[i] {
			stack = stack[:len(stack)-1]
			i++
		}
	}
	//总结判断
	return len(stack) == 0
}

func main() {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 2, 1}
	res := validateStackSequences(pushed, popped)
	fmt.Println(res)
}
