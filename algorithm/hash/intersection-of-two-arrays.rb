# // 349. 两个数组的交集
# 力扣题目链接(https://leetcode-cn.com/problems/happy-number/)
# // 给定两个数组，编写一个函数来计算它们的交集。

# // 示例 1：
# // 输入：nums1 = [1,2,2,1], nums2 = [2,2]
# // 输出：[2]

# // 示例 2：
# // 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
# // 输出：[9,4]

# // 说明：

# // 输出结果中的每个元素一定是唯一的。
# // 我们可以不考虑输出结果的顺序。

def intersection(nums1, nums2)
	set = {}
	res = []
	nums1.each do |num|
		set[num] = true
	end
	
	nums2.each do |num|
		#如果存在于上一个数组中，则加入结果集，并清空该set值
		if set[num]
			res << num
			set.delete num
		end
	end
	
	res
end

# 输入：nums1 = [1,2,2,1], nums2 = [2,2]
# 输出：[2]
nums1 = [1, 2, 2, 1]
nums2 = [2, 2]
res = intersection(nums1, nums2)
p res # should be [2]

# 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
# 输出：[9,4]
nums1 = [4, 9, 5]
nums2 = [9, 4, 9, 8, 4]
res = intersection(nums1, nums2)
p res # should be [9,4]
