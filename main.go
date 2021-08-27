package main

import (
	"fmt"
)

func AddElement(slice []int, e int) []int {
	return append(slice, e)
}

func main() {
	var slice []int
	fmt.Println(cap(slice))
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	fmt.Println(cap(slice))

	newSlice := AddElement(slice, 4)
	fmt.Println(cap(newSlice))
	fmt.Println(&(slice[0]))
	fmt.Println(&(newSlice[0]))
	fmt.Println(&slice[0] == &newSlice[0])
}

// func main() {
// 	arr := [...]int{1, 2, 3}
// 	fmt.Println(len(arr))

// 	fmt.Println(10 % 2)
// 	fmt.Println(11 % 2)
// 	fmt.Println(math.Pow(3, 2))
// 	fmt.Println(5 / 2)
// 	fmt.Println(int64(5) / 2)
// 	fmt.Println(float64(5) / 2)
// 	fmt.Println(float32(5) / 2)

// 	var slice []int
// 	fmt.Println(cap(slice))
// 	slice = append(slice, 1)
// 	fmt.Println(cap(slice))
// 	slice = append(slice, 2)
// 	fmt.Println(cap(slice))
// 	slice = append(slice, 3)
// 	fmt.Println(cap(slice))

// }
