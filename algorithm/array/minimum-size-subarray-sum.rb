# // 209.长度最小的子数组
# // 力扣题目链接(https://leetcode-cn.com/problems/minimum-size-subarray-sum/)
#
# // 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。
#
# // 示例：
#
# // 输入：s = 7, nums = [2,3,1,2,4,3] 输出：2 解释：子数组 [4,3] 是该条件下的长度最小的子数组。

def min_sub_array_len(target, nums)
	i = 0               # 滑动窗口起始位置
	length = nums.length  # 滑动窗口的长度
	sum = 0             # 滑动窗口数值之和
	result = length + 1 # 初始化返回长度为length+1，目的是为了判断“不存在符合条件的子数组，返回0”的情况

	nums.each_index do |j|
		sum += nums[j]
		# 注意这里使用while，每次更新 i（起始位置），并不断比较子序列是否符合条件
		while sum >= target
			sub_length = j - i + 1 # 取子序列的长度
			if sub_length < result
				result = sub_length
			end
			sum -= nums[i]   # 这里体现出滑动窗口的精髓之处，不断变更i（子序列的起始位置）,和现在需要减去nums[i]
			i += 1           # 这里体现出滑动窗口的精髓之处，不断变更i（子序列的起始位置）
		end
	end

	# 如果result没有被赋值的话，就返回0，说明没有符合条件的子序列
	if result == length+1
		return 0
	else
		return result
	end
end

target = 7
nums = [2, 3, 1, 2, 4, 3]
res = min_sub_array_len(target, nums)
p res
