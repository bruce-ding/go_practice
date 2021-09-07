package main

import (
	"time"
)

func main() {
	var c1 chan string = make(chan string)
	func() {
		time.Sleep(time.Second)
		c1 <- "1"
		//fmt.Println("c1 is", <-c1)
	}()
	//c1 <- "1"

	//fmt.Println("c1 is", <-c1)
}
