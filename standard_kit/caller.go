package main

import (
	"fmt"
	"runtime"
)

func main() {
	var ok bool = true
	var i int
	for ok {
		ok = test(i)
		i++
	}
}

func test(skip int) bool {
	return call(skip)
}

func call(skip int) bool {
	pc, file, line, ok := runtime.Caller(skip)
	pcName := runtime.FuncForPC(pc).Name() //获取函数名
	if ok {
		fmt.Printf("%v\t%s\t%d\t%t\t%s\n", pc, file, line, ok, pcName)
	}

	return ok
}
