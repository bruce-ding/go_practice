package main

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}

	h := &MinHeap{}
	heap.Init(h)
	for key, value := range occurrences {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

type MinHeap [][2]int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// // An Item is something we manage in a priority queue.
// type Item struct {
// 	value    string // The value of the item; arbitrary.
// 	priority int    // The priority of the item in the queue.
// 	// The index is needed by update and is maintained by the heap.Interface methods.
// 	index int // The index of the item in the heap.
// }

// // A PriorityQueue implements heap.Interface and holds Items.
// type PriorityQueue []*Item

// func (pq PriorityQueue) Len() int { return len(pq) }

// func (pq PriorityQueue) Less(i, j int) bool {
// 	return pq[i].priority < pq[j].priority
// }

// func (pq PriorityQueue) Swap(i, j int) {
// 	pq[i], pq[j] = pq[j], pq[i]
// 	pq[i].index = i
// 	pq[j].index = j
// }

// func (pq *PriorityQueue) Push(x interface{}) {
// 	n := len(*pq)
// 	item := x.(*Item)
// 	item.index = n
// 	*pq = append(*pq, item)
// }

// func (pq *PriorityQueue) Pop() interface{} {
// 	old := *pq
// 	n := len(old)
// 	item := old[n-1]
// 	old[n-1] = nil  // avoid memory leak
// 	item.index = -1 // for safety
// 	*pq = old[0 : n-1]
// 	return item
// }

// // update modifies the priority and value of an Item in the queue.
// func (pq *PriorityQueue) update(item *Item, value string, priority int) {
// 	item.value = value
// 	item.priority = priority
// 	heap.Fix(pq, item.index)
// }

func main() {

}
