package stack_queue

/**
347.前 K 个高频元素
力扣题目链接(https://leetcode-cn.com/problems/top-k-frequent-elements/)

给定一个非空的整数数组，返回其中出现频率前 k 高的元素。

示例 1:
输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]

示例 2:
输入: nums = [1], k = 1
输出: [1]

提示：
你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。
题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的。
你可以按任意顺序返回答案。
**/

import (
	"container/heap"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 方法一：小顶堆 时间复杂度 O(Nlogk)
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
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return res
}

type MinHeap [][2]int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

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

// 方法二: 利用O(logn)排序
func topKFrequent1(nums []int, k int) []int {
	res := []int{}
	m := map[int]int{}
	for _, item := range nums {
		m[item]++
	}
	for key := range m {
		res = append(res, key)
	}
	//核心思想：排序
	//可以不用包函数，自己实现快排
	sort.Slice(res, func(i, j int) bool {
		return m[res[i]] > m[res[j]]
	})

	return res[:k]
}
func TestTopKFrequent(t *testing.T) {
	// 	示例 1:
	// 输入: nums = [1,1,1,2,2,3], k = 2
	// 输出: [1,2]
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	res := topKFrequent(nums, k)
	assert.Equal(t, []int{1, 2}, res)

	res = topKFrequent1(nums, k)
	assert.Equal(t, []int{1, 2}, res)

	// 示例 2:
	// 输入: nums = [1], k = 1
	// 输出: [1]
	nums = []int{1}
	k = 1
	res = topKFrequent(nums, k)
	assert.Equal(t, []int{1}, res)

	res = topKFrequent1(nums, k)
	assert.Equal(t, []int{1}, res)
}
