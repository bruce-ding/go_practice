package main

import (
	"fmt"
	"golang.org/x/time/rate"
)

func main() {
	//limit := rate.Every(100 * time.Millisecond);
	limiter := rate.NewLimiter(10, 1)
	limiter.Allow()
	b := limiter.Burst()
	fmt.Println(b)
}
