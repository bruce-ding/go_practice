package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func main() {
	info, _ := cpu.Info()

	per, _ := cpu.Percent(1*time.Second, true)

	fmt.Printf("CPU Percent: %f\n", per)

	fmt.Println(info)
}
