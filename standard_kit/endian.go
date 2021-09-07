package main

import (
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"unsafe"
)


func littleEndianUint32(b []byte) uint32 {
	_ = b[3] // bounds check hint to compiler; see golang.org/issue/14808
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

func littleEndianPutUint32(b []byte, v uint32) {
	_ = b[3] // early bounds check to guarantee safety of writes below
	fmt.Println(v)
	fmt.Println(byte(v))

	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)

}

func littleEndianUint64(b []byte) uint64 {
	_ = b[7] // bounds check hint to compiler; see golang.org/issue/14808
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
}

func littleEndianPutUint64(b []byte, v uint64) {
	_ = b[7] // early bounds check to guarantee safety of writes below
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	b[4] = byte(v >> 32)
	b[5] = byte(v >> 40)
	b[6] = byte(v >> 48)
	b[7] = byte(v >> 56)
}

const IntSize int = int(unsafe.Sizeof(0))

//判断我们系统中的字节序类型
func systemEndian() {
	var i int = 0x1
	bs := (*[IntSize]byte)(unsafe.Pointer(&i))
	fmt.Println(len(bs))
	if bs[0] == 0 {
		fmt.Println("system endian is little endian")
	} else {
		fmt.Println("system endian is big endian")
	}
}

func testBigEndian() {
	// 0000 0000 0000 0000   0000 0001 1111 1111
	var testInt int32 = 256
	fmt.Printf("%d use big endian: \n", testInt)
	var testBytes []byte = make([]byte, 4)
	binary.BigEndian.PutUint32(testBytes, uint32(testInt))
	fmt.Println("int32 to bytes:", testBytes)

	convInt := binary.BigEndian.Uint32(testBytes)
	fmt.Printf("bytes to int32: %d\n\n", convInt)
}

func testLittleEndian() {
	// 0000 0000 0000 0000   0000 0001 1111 1111
	var testInt int32 = 256
	fmt.Printf("%d use little endian: \n", testInt)
	var testBytes []byte = make([]byte, 4)
	//binary.LittleEndian.PutUint32(testBytes, uint32(testInt))
	littleEndianPutUint32(testBytes, uint32(testInt))
	fmt.Println("int32 to bytes:", testBytes)


	convInt := binary.LittleEndian.Uint32(testBytes)
	fmt.Printf("bytes to int32: %d\n\n", convInt)
}

func main() {
	systemEndian()
	fmt.Println("")
	testBigEndian()
	testLittleEndian()
	//
	//fmt.Println(byte(100))
	//fmt.Println(256 | 1 | 0 | 0)
	//fmt.Println(256 >> 24)

	//fmt.Println(byte(uint32(256)))

	fmt.Println("slice")
	slice := make([]int, 0, 5)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	fmt.Println(slice)
	slice = append(slice, 1)
	fmt.Println(slice)
	slice = append(slice, 2)
	fmt.Println(slice)
	slice = append(slice, 3, 4)
	fmt.Println(slice)
	fmt.Println("trim")

	//fmt.Println(slice[1:1])
	//fmt.Println(slice[2:3]) // 1 2 3 4
	//fmt.Println(slice[0:3])
	//fmt.Println(slice[0:4])

	slice01 := []int{1,2,3,4}
	var slice02 = []int{7}
	copy(slice01[1:2], slice02)
	fmt.Println(slice01[1])

	slice01 = append(slice01,  slice02...)

	fmt.Println("copy")
	fmt.Println(slice01)

	//copy(slice01, e.Meta.Key)

	crc := crc32.ChecksumIEEE([]byte("hello world"))
	fmt.Println(crc)

	m := make(map[string]string)
	m["123"] = "hello"
	v, ok := m["123"]
	if ok {
		fmt.Println("key 123:", v)
	}
}
