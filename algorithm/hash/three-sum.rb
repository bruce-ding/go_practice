# // 15. 三数之和
# // 力扣题目链接(https://leetcode-cn.com/problems/3sum/)
# // 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

# // 注意：答案中不可以包含重复的三元组。

# // 示例 1：
# // 输入：nums = [-1,0,1,2,-1,-4]
# // 输出：[[-1,-1,2],[-1,0,1]]

# // 示例 2：
# // 输入：nums = []
# // 输出：[]

# // 示例 3：
# // 输入：nums = [0]
# // 输出：[]

# // 提示：
# // 0 <= nums.length <= 3000
# // -105 <= nums[i] <= 105


# 排序 + 双指针
# 本题的难点在于如何去除重复解。

# 算法流程：
# 特判，对于数组长度 n，如果数组长度小于 3，返回 []。
# 对数组进行排序。
# 遍历排序后数组：
# 若 nums[i]>0：因为已经排序好，所以后面不可能有三个数加和等于 0，直接返回结果。
# 对于重复元素：跳过，避免出现重复解
# 令左指针 L=i+1，右指针 R=n-1，当 L<R 时，执行循环：
# 当 nums[i]+nums[L]+nums[R]==0，执行循环，判断左界和右界是否和下一位置重复，去除重复解。并同时将 L,R 移到下一位置，寻找新的解
# 若和大于 0，说明 nums[R] 太大，R 左移
# 若和小于 0，说明 nums[L] 太小，L 右移


def three_sum(nums, target)
	res = []
	n = nums.length
	# 对于数组长度n，如果数组长度小于3，直接返回
	if n < 3
		return res
	end
	nums.sort! # O(nlgn)

	nums.each_index do |i|
		# 若 nums[i]>0：因为已经排序好，所以后面不可能有三个数加和等于 0，直接返回结果
		if nums[i] > target
			return res
		end

		# 对于重复元素：跳过，避免出现重复解
		if i > 0 && nums[i] == nums[i-1]
			next
		end

		left = i + 1
		right = n - 1
		while left < right
			sum = nums[i] + nums[left] + nums[right]
			if sum == target
				res << [nums[i], nums[left], nums[right]]

				# 判断左界和右界是否和下一位置重复，去除重复解
				while left < right && nums[left] == nums[left+1]
					left += 1
				end
				while left < right && nums[right] == nums[right-1]
					right -= 1
				end

				left += 1
				right -= 1
			elsif sum > target
				right -= 1
			else
				left += 1
			end
		end
	end
	
	res
end

# 输入：nums = [-1,0,1,2,-1,-4]
# 输出：[[-1,-1,2],[-1,0,1]]
nums = [-1, 0, 1, 2, -1, -4]
target = 0
res = three_sum(nums, target)
p res # should be [[-1, -1, 2], [-1, 0, 1]]