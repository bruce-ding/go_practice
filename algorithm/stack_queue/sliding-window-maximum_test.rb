=begin
239. 滑动窗口最大值
力扣题目链接(https:#leetcode-cn.com/problems/sliding-window-maximum/)
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
=end

# 封装单调(递减)队列的方式解题
class MonotonicQueue
	attr_accessor :queue

	def initialize()
	  @queue = []
	end

	def front() # 第一个元素，也是最大值
		return @queue.first
	end
	
	def back() # 最后一个元素，也是最小值
		return @queue.last
	end
	
	def empty()
		return @queue.length == 0
	end
	
	def push(val)
		while !self.empty() && val > self.back() 
			@queue.pop
		end
		@queue.push val
	end
	
	def pop(val)
		if !self.empty() && val == self.front() 
			@queue.shift # remove first element and return its value
		end
	end
end

def max_sliding_window(nums, k)
	mq = MonotonicQueue.new
	length = nums.length
	res = []

	# 先将前k个元素放入队列
	k.times do |i|
		mq.push(nums[i])
	end
	# 记录前k个元素的最大值
	res.push mq.front

	(k..length-1).each do |i|
		# 滑动窗口移除最前面的元素
		mq.pop(nums[i-k])
		# 滑动窗口添加最后面的元素
		mq.push(nums[i])
		# 记录最大值
		res.push mq.front()
	 end
	return res
end

require 'minitest/autorun'

describe "max_sliding_window" do
	it "must" do
		#输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
		#输出：[3,3,5,5,6,7]
		#解释：
		# 滑动窗口的位置                 最大值
		#---------------               -----
		#[1  3  -1] -3  5  3  6  7       3
		#1  [3  -1  -3] 5  3  6  7       3
		#1   3 [-1  -3  5] 3  6  7       5
		#1   3  -1 [-3  5  3] 6  7       5
		#1   3  -1  -3 [5  3  6] 7       6
		#1   3  -1  -3  5 [3  6  7]      7
		nums = [1, 3, -1, -3, 5, 3, 6, 7]
		k = 3
		res = max_sliding_window(nums, k)
		assert_equal([3, 3, 5, 5, 6, 7], res)
	end
end
 