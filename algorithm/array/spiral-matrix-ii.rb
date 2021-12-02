# //59.螺旋矩阵II
# //力扣题目链接(https://leetcode-cn.com/problems/spiral-matrix-ii/)
# //
# //给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。
# //
# //示例:
# //
# //输入: 3 输出: [ [ 1, 2, 3 ], [ 8, 9, 4 ], [ 7, 6, 5 ] ]

def spiral_matrix(n)
	matrix = Array.new(n)
	matrix.each_index do |i|
		matrix[i] = Array.new(n)
	end

	left, right, top, bottom = 0, n-1, 0, n-1
	size = n * n
	step = 0

	while step < size
		i = left
		while i <= right && step < size
			matrix[top][i] = step + 1
			step += 1

			i += 1
		end
		top += 1

		i = top
		while i <= bottom && step < size
			matrix[i][right] = step + 1
			step += 1

			i += 1
		end
		right -= 1

		i = right
		while i >= left && step < size
			matrix[bottom][i] = step + 1
			step += 1

			i -= 1
		end
		bottom -= 1

		i = bottom
		while i >= top && step < size
			matrix[i][left] = step + 1
			step += 1

			i -= 1
		end
		left += 1
	end

	return matrix
end

# 1  2  3
# 8  9  4
# 7  6  5
matrix = spiral_matrix(3)
p matrix
