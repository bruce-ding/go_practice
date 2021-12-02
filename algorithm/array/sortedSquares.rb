# // 977.有序数组的平方
# // 力扣题目链接(https://leetcode-cn.com/problems/squares-of-a-sorted-array/)

# // 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

# // 示例 1： 输入：nums = [-4,-1,0,3,10] 输出：[0,1,9,16,100] 解释：平方后，数组变为 [16,1,0,9,100]，排序后，数组变为 [0,1,9,16,100]

# // 示例 2： 输入：nums = [-7,-3,2,3,11] 输出：[4,9,9,49,121]

def sorted_squares(nums)
	n = nums.length
	i, j, k = 0, n-1, n-1
	ans = Array.new(n)

	while i <= j
		lm, rm = nums[i]**2, nums[j]**2
		if lm > rm
			ans[k] = lm
			i += 1
		else
			ans[k] = rm
			j -= 1
		end
		k -= 1
	end

	ans
end

# nums = [-4,-1,0,3,10] 输出：[0,1,9,16,100] 解释：平方后，数组变为 [16,1,0,9,100]，排序后，数组变为 [0,1,9,16,100]

nums = [-4, -1, 0, 3, 10]
res = sorted_squares(nums)
p res
