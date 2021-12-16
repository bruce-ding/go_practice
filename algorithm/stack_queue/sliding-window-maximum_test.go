package stack_queue

/**
239. 滑动窗口最大值
力扣题目链接(https://leetcode-cn.com/problems/sliding-window-maximum/)
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
返回滑动窗口中的最大值。

示例 1：
输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]

解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7

示例 2：
输入：nums = [1], k = 1
输出：[1]

示例 3：
输入：nums = [1,-1], k = 1
输出：[1,-1]

示例 4：
输入：nums = [9,11], k = 2
输出：[11]

示例 5：
输入：nums = [4,-2], k = 2
输出：[4]

提示：
1 <= nums.length <= 105
-104 <= nums[i] <= 104
1 <= k <= nums.length
**/

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

// 封装单调队列的方式解题
type MonotonicQueue struct {
	queue []int
}

func NewMonotonicQueue() *MonotonicQueue {
	return &MonotonicQueue{
		queue: make([]int, 0),
	}
}

func (m *MonotonicQueue) Front() int {
	return m.queue[0]
}

func (m *MonotonicQueue) Back() int {
	return m.queue[len(m.queue)-1]
}

func (m *MonotonicQueue) Empty() bool {
	return len(m.queue) == 0
}

func (m *MonotonicQueue) Push(val int) {
	for !m.Empty() && val > m.Back() {
		m.queue = m.queue[:len(m.queue)-1]
	}
	m.queue = append(m.queue, val)
}

func (m *MonotonicQueue) Pop(val int) {
	if !m.Empty() && val == m.Front() {
		m.queue = m.queue[1:]
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	queue := NewMonotonicQueue()
	length := len(nums)
	res := make([]int, 0)
	// 先将前k个元素放入队列
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
		spew.Dump(queue)
	}
	// 记录前k个元素的最大值
	res = append(res, queue.Front())

	for i := k; i < length; i++ {
		// 滑动窗口移除最前面的元素
		queue.Pop(nums[i-k])
		// 滑动窗口添加最后面的元素
		queue.Push(nums[i])
		spew.Dump(queue)
		// 记录最大值
		res = append(res, queue.Front())
	}
	return res
}

func TestMaxSlidingWindow(t *testing.T) {
	//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
	//输出：[3,3,5,5,6,7]
	//解释：
	//滑动窗口的位置                最大值
	//---------------               -----
	//[1  3  -1] -3  5  3  6  7       3
	//1 [3  -1  -3] 5  3  6  7       3
	//1  3 [-1  -3  5] 3  6  7       5
	//1  3  -1 [-3  5  3] 6  7       5
	//1  3  -1  -3 [5  3  6] 7       6
	//1  3  -1  -3  5 [3  6  7]      7
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	res := maxSlidingWindow(nums, k)
	assert.Equal(t, []int{3, 3, 5, 5, 6, 7}, res)
}
