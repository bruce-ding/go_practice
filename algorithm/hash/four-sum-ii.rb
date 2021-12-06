# // 454. 四数相加 II
# // 力扣题目链接(https://leetcode-cn.com/problems/4sum-ii/)
# // 给你四个整数数组 nums1、nums2、nums3 和 nums4 ，数组长度都是 n ，请你计算有多少个元组 (i, j, k, l) 能满足：

# // 0 <= i, j, k, l < n
# // nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0

# // 示例 1：
# // 输入：nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2]
# // 输出：2
# // 解释：
# // 两个元组如下：
# // 1. (0, 0, 0, 1) -> nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) + (-1) + 2 = 0
# // 2. (1, 1, 0, 0) -> nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) + (-1) + 0 = 0

# // 示例 2：
# // 输入：nums1 = [0], nums2 = [0], nums3 = [0], nums4 = [0]
# // 输出：1

def four_sum_count(nums1, nums2, nums3, nums4)
	m = {}
	count = 0
	nums1.each do |v1|
		nums2.each do |v2|
			if m.has_key?(v1+v2)
				m[v1+v2] += 1
			else
				m[v1+v2] = 1
			end
		end
	end
	nums3.each do |v3|
		nums4.each do |v4|
			count += m[-v3-v4].to_i
		end
	end
	
	count
end

# 输入：nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2]
# 输出：2
# 解释：
# 两个元组如下：
# 1. (0, 0, 0, 1) -> nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) + (-1) + 2 = 0
# 2. (1, 1, 0, 0) -> nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) + (-1) + 0 = 0
nums1 = [1, 2]
nums2 = [-2, -1]
nums3 = [-1, 2]
nums4 = [0, 2]
res = four_sum_count(nums1, nums2, nums3, nums4)
p(res) # should be 2
