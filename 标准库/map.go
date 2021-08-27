package main

import "unsafe"

type hchan1 struct {
	qcount   uint           // 当前队列中剩余元素个数
	dataqsiz uint           // 环形队列长度，即可以存放的元素个数
	buf      unsafe.Pointer // 环形队列指针
	elemsize uint16         // 每个元素的大小
	closed   uint32         // 标识关闭状态
	elemtype *_type         // 元素类型
	sendx    uint           // 队列下标，指示元素写入时存放到队列中的位置
	recvx    uint           // 队列下标，指示元素从队列的该位置读出
	recvq    waitq          // 等待读消息的goroutine队列
	sendq    waitq          // 等待写消息的goroutine队列
	loc
}

type hchan struct {
	qcount uint // 队列元素个数
	dataqsiz uint// 队列长度，即可以存放的元素个数
	buf unsafe.Pointer // 环形队列指针
	elemsize uint16 // 每个元素的大小
	closed uint32 // 标识关闭状态
	elemtype *_type // 元素类型
	sendx uint
	recvx uint
	recvq waitq
	sendq waitq
	loc mutex
}

func main() {

}
