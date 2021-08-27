package main

import "fmt"

// 遍历添加
func replaceSpace1(s string) string {
	b := []byte(s)
	result := make([]byte, 0)
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' {
			result = append(result, []byte("%20")...)
		} else {
			result = append(result, b[i])
		}
	}
	return string(result)
}

// func replaceSpace2(s string) string {
// 	bytes := []byte(s)
// 	l := len(bytes)
// 	count := 0
// 	for i := 0; i < l; i++ {
// 		if bytes[i] == ' ' {
// 			count++
// 		}
// 	}
// 	fmt.Println("count", count)
// 	result := make([]byte, l+count*2)
// 	fmt.Printf("cap: %v len: %v\n", cap(result), len(result))
// 	j := 0
// 	for i := 0; i < l; i++ {
// 		if bytes[i] == ' ' {
// 			result[j] = '%'
// 			result[j+1] = '2'
// 			result[j+2] = '0'
// 			j += 2
// 		} else {
// 			result[j] = bytes[j]
// 			j++
// 		}
// 	}
// 	fmt.Println("result", result)

// 	return string(result)
// }

// 原地修改
func replaceSpace(s string) string {
	b := []byte(s)
	length := len(b)
	spaceCount := 0
	// 计算空格数量
	for _, v := range b {
		if v == ' ' {
			spaceCount++
		}
	}
	// 扩展原有切片
	resizeCount := spaceCount * 2
	tmp := make([]byte, resizeCount)
	b = append(b, tmp...)

	i := length - 1
	j := len(b) - 1
	for i >= 0 {
		if b[i] != ' ' {
			b[j] = b[i]
			i--
			j--
		} else {
			b[j] = '0'
			b[j-1] = '2'
			b[j-2] = '%'
			i--
			j -= 3
		}
	}

	return string(b)
}

func main() {
	str := replaceSpace("We are happy.")
	fmt.Println(str)
	fmt.Println(len([]byte(" ")))
}
