package main

import "fmt"

func sortArray(nums []int) []int {
	return quicksort(nums, 0, len(nums)-1)
}

func quicksort(nums []int, low int, high int) []int {
	if low >= high {
		return nums
	}
	pivot := nums[low]
	start := low
	end := high
	for low < high {
		for low < high && nums[high] >= pivot {
			high--
		}
		if low < high {
			nums[low] = nums[high]
			low++
		}
		for low < high && nums[low] < pivot {
			low++
		}
		if low < high {
			nums[high] = nums[low]
			high--
		}

	}
	nums[low] = pivot
	quicksort(nums, start, low-1)
	quicksort(nums, low+1, end)
	return nums
}

func main() {
	arr := []int{1, 2, 6, 4, 8}
	res := sortArray(arr)
	fmt.Println(res)
}
