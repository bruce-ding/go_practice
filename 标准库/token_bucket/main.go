package main

import (
	"fmt"
	"time"
)

func main() {
	l := newLimiter(60, time.Second)
	ticker := time.NewTicker(time.Second / 2)

	for {
		select {
		case <-ticker.C:
			b := l.allow()
			fmt.Println("allow:", b)
		}
	}

}
