package leetcode

func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	start := left
	left, right = 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= target+1 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	end := left - 1
	return []int{start, end}
}
