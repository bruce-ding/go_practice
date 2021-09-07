package main

import (
	"fmt"
	"runtime"
)

func DumpProcessTrace() {
	buf := make([]byte, 64*1024)
	_ = runtime.Stack(buf, true)
	fmt.Println("FULL PROCESS THREAD DUMP:")
	fmt.Println(string(buf))
}

func demo() {
	DumpProcessTrace()
}

func main() {
	demo()
}
