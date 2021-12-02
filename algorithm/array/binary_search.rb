# //704. 二分查找
# //力扣题目链接(https://leetcode-cn.com/problems/binary-search/)
# //
# //给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
# //
# //示例 1:
# //
# //输入: nums = [-1,0,3,5,9,12], target = 9
# //输出: 4
# //解释: 9 出现在 nums 中并且下标为 4
# //示例 2:
# //
# //输入: nums = [-1,0,3,5,9,12], target = 2
# //输出: -1
# //解释: 2 不存在 nums 中因此返回 -1
# //提示：
# //
# //你可以假设 nums 中的所有元素是不重复的。
# //n 将在 [1, 10000]之间。
# //nums 的每个元素都将在 [-9999, 9999]之间。

def binarySearch(nums, target)
    left = 0
    right = nums.length - 1
    while left <= right
        middle = left + (right - left) / 2
        if nums[middle] > target
            right = middle - 1
        elsif nums[middle] < target
            left = middle + 1
        else
            return middle
        end
    end

    -1
end

nums = [1,2,3,4,5,6,7,8,9]
res = binarySearch(nums, 8)
p res
